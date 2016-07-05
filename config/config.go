package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	configPrefix = "paus"
)

type Config struct {
	DatadogAPIKey string `envconfig:"datadog_api_key" default:""`
	DatadogAppKey string `envconfig:"datadog_app_key" default:""`
	EtcdEndpoint  string `envconfig:"etcd_endpoint" default:"http://localhost:2379"`
	TargetKey     string `envconfig:"target_key" required:"true"`
}

func LoadConfig() (*Config, error) {
	var config Config

	err := envconfig.Process(configPrefix, &config)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}

	return &config, nil
}
