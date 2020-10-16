package hooks

import (
	sentry "github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type SentryHook struct{}

func NewSentryHook() *SentryHook {
	return new(SentryHook)
}

func (h SentryHook) Run(_ *zerolog.Event, level zerolog.Level, msg string) {
	if level >= zerolog.ErrorLevel {
		sentry.CaptureMessage(msg)
	}
}
