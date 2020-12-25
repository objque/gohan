package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig_LoadFromFile(t *testing.T) {
	// arrange
	conf := New()
	expected := AppConfig{
		Log: LogConfig{
			Level: "DEBUG",
		},
		HTTP: HTTPConfig{
			IP:           "0.0.0.0",
			Port:         8080,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
		Sentry: SentryConfig{
			Enabled:      false,
			DSN:          "https://xxxx@sentry.org/100",
			Environment:  "development",
			FlushTimeout: 10 * time.Second,
		},
	}

	// action
	err := conf.LoadFromFile("../../gohan.example.yml")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, *conf)
}
