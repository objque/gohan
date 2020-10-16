package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/objque/gohan/internal/log"
	"github.com/objque/gohan/internal/log/hooks"
	"github.com/objque/gohan/internal/version"
)

func main() {
	log.SetLevel("DEBUG")
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

	log.Debug(version.FullInfo)
	log.Info("Hello, world")
}

func exitIfError(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(2) //nolint:gomnd
}
