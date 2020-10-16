package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/objque/go-app-template/internal/log"
	"github.com/objque/go-app-template/internal/log/hooks"
)

func main() {
	log.SetLevel("INFO")
	log.SetWriters(log.GetConsoleWriter())

	if os.Getenv("GOHAN_SENTRY_DSN") != "" {
		err := sentry.Init(sentry.ClientOptions{
			AttachStacktrace: true,
			Dsn:              os.Getenv("GOHAN_SENTRY_DSN"),
			Environment:      os.Getenv("GOHAN_SENTRY_ENVIRONMENT"),
			Release:          "1.0",
		})
		if err != nil {
			exitIfError(fmt.Errorf("sentry initialization failed: %w", err))
		}

		log.SetHook(hooks.NewSentryHook())
		log.Info("sentry integration enabled")
	}

	log.Info("Hello, world")
}

func exitIfError(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
