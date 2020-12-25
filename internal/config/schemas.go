package config

import "time"

type AppConfig struct {
	Log    LogConfig    `yaml:"log"`
	HTTP   HTTPConfig   `yaml:"http"`
	Sentry SentryConfig `yaml:"sentry"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type HTTPConfig struct {
	IP           string        `yaml:"ip"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type SentryConfig struct {
	Enabled      bool          `yaml:"enabled"`
	DSN          string        `yaml:"dsn"`
	Environment  string        `yaml:"environment"`
	FlushTimeout time.Duration `yaml:"flush_timeout"`
}
