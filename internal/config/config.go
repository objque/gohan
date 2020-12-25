package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func New() *AppConfig {
	return &AppConfig{
		Log: LogConfig{
			Level: "DEBUG",
		},
		HTTP: HTTPConfig{
			IP:           "127.0.0.1",
			Port:         7711,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
		Sentry: SentryConfig{
			FlushTimeout: 10 * time.Second,
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
	return string(b)
}
