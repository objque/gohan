package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/getsentry/sentry-go"
	"github.com/objque/gohan/internal/api"
	"github.com/objque/gohan/internal/config"
	"github.com/objque/gohan/internal/log"
	"github.com/objque/gohan/internal/log/hooks"
	"github.com/objque/gohan/internal/services/subscriptions"
	"github.com/objque/gohan/internal/version"
)

func main() {
	configPath := flag.String("config", "", "Abs path to config.yaml")

	flag.Parse()

	conf := config.New()
	if *configPath != "" {
		exitIfError(conf.LoadFromFile(*configPath))
	}

	log.SetLevel(conf.Log.Level)
	log.SetWriters(log.GetConsoleWriter())
	log.Debug(conf.Dump())

	if conf.Sentry.Enabled {
		err := sentry.Init(sentry.ClientOptions{
			AttachStacktrace: true,
			Dsn:              conf.Sentry.DSN,
			Environment:      conf.Sentry.Environment,
			Release:          version.Commit,
		})
		if err != nil {
			exitIfError(fmt.Errorf("sentry initialization failed: %w", err))
		}

		log.SetHook(hooks.NewSentryHook())
		log.Info("sentry integration enabled")
	}

	routerOpts := api.RouterOpts{
		SubscriptionsRepository: subscriptions.New(),
	}
	server := api.New(api.GetRouter(routerOpts), conf.HTTP)

	log.Debug("gohan started")
	log.Info(version.FullInfo)

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithTimeout(context.Background(), conf.HTTP.WriteTimeout)
	defer cancel()

	go gracefulShutdown(ctx, server, quit, done)

	log.Infof("server is ready to handle requests at: %v", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		exitIfError(fmt.Errorf("could not listen on %v: %w", server.Addr, err))
	}

	<-done
	log.Info("gohan stopped")
}

func gracefulShutdown(ctx context.Context, server *api.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	log.Info("server is shutting down...")

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("could not gracefully shutdown the server: %v", err)
	}
	close(done)
}

func exitIfError(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(2) //nolint:gomnd
}
