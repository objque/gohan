package log

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	callerSkipFrameCount = 3
)

//nolint:gochecknoglobals
var (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel

	logger = zerolog.New(os.Stderr).With().CallerWithSkipFrameCount(callerSkipFrameCount).Timestamp().Logger()
)

func GetConsoleWriter() io.Writer {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02T15:04:05",
	}
}

func SetHook(h zerolog.Hook) {
	logger = logger.Hook(h)
}

func SetWriters(writers ...io.Writer) {
	logger = logger.Output(zerolog.MultiLevelWriter(writers...))
}

func SetLevel(lvl string) {
	switch strings.ToLower(lvl) {
	case "trace":
		logger = logger.Level(zerolog.TraceLevel)
	case "debug":
		logger = logger.Level(zerolog.DebugLevel)
	case "info":
		logger = logger.Level(zerolog.InfoLevel)
	case "warn":
		logger = logger.Level(zerolog.WarnLevel)
	case "error":
		logger = logger.Level(zerolog.ErrorLevel)
	case "fatal":
		logger = logger.Level(zerolog.FatalLevel)
	case "panic":
		logger = logger.Level(zerolog.PanicLevel)
	}
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug().Msgf(format, args...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Infof(format string, args ...interface{}) {
	logger.Info().Msgf(format, args...)
}

func Error(msg string) {
	logger.Error().Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	logger.Error().Msgf(format, args...)
}

func Warn(msg string) {
	logger.Warn().Msg(msg)
}

func Warnf(format string, args ...interface{}) {
	logger.Warn().Msgf(format, args...)
}
