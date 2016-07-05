package provider

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/zorkian/go-datadog-api"
)

type Datadog struct {
	client *datadog.Client
}

func NewDatadog(apiKey, appKey string) Datadog {
	return Datadog{
		client: datadog.NewClient(apiKey, appKey),
	}
}

func (d Datadog) Notify(action, key, value string) error {
	if action != "set" {
		return nil
	}

	base := filepath.Base(key)
	timestamp, err := strconv.Atoi(base)

	if err != nil {
		return errors.Wrapf(err, "Failed to parse timestamp: %s", base)
	}

	event := &datadog.Event{
		Title: fmt.Sprintf("Deployed new revision: %s", value),
		Text:  key,
		Time:  timestamp,
		Tags:  []string{"paus:deployment"},
	}

	_, err = d.client.PostEvent(event)

	if err != nil {
		return errors.Wrap(err, "Failed to post event")
	}

	return nil
}
