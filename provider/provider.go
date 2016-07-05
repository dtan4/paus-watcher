package provider

import (
	"github.com/dtan4/paus-watcher/config"
)

type Provider interface {
	Notify(action, key, value string) error
}

func NewProvider(config *config.Config) Provider {
	if (config.DatadogAPIKey != "") && (config.DatadogAppKey != "") {
		return NewDatadog(config.DatadogAPIKey, config.DatadogAppKey)
	}

	return nil
}
