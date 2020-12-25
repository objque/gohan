package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	// log
	defaultLogLevel = "debug"
	// http
	defaultIP           = "127.0.0.1"
	defaultPort         = 7711
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
	defaultIdleTimeout  = 30 * time.Second

	// sentry
	defaultFlushTimeout = 10 * time.Second
)

func New() *AppConfig {
	return &AppConfig{
		Log: LogConfig{
			Level: defaultLogLevel,
		},
		HTTP: HTTPConfig{
			IP:           defaultIP,
			Port:         defaultPort,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			IdleTimeout:  defaultIdleTimeout,
		},
		Sentry: SentryConfig{
			FlushTimeout: defaultFlushTimeout,
		},
	}
}

func (c *AppConfig) LoadFromFile(configPath string) error {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	// replace ${ENV_NAME} in file with value from the environment
	b = []byte(os.ExpandEnv(string(b)))

	return c.LoadFromBytes(b)
}

func (c *AppConfig) LoadFromBytes(val []byte) error {
	return yaml.Unmarshal(val, c)
}

func (c *AppConfig) Dump() string {
	b, _ := yaml.Marshal(c)
	return string(b) //nolint:nlreturn
}
