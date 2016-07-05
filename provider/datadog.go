package provider

import (
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

	timestamp, err := strconv.Atoi(value)

	if err != nil {
		return errors.Wrapf(err, "Failed to parse timestamp: %s", value)
	}

	event := &datadog.Event{
		Title: "Deployed new revision",
		Text:  key,
		Time:  timestamp,
	}

	_, err = d.client.PostEvent(event)

	if err != nil {
		return errors.Wrap(err, "Failed to post event")
	}

	return nil
}
