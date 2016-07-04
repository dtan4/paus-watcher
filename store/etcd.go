package store

import (
	"github.com/coreos/etcd/client"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Etcd struct {
	keysAPI client.KeysAPI
}

type EtcdWatchFunc func(*client.Response)

func NewEtcd(endpoint string) (*Etcd, error) {
	config := client.Config{
		Endpoints: []string{endpoint},
		Transport: client.DefaultTransport,
	}

	c, err := client.New(config)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to create etcd client.")
	}

	keysAPI := client.NewKeysAPI(c)

	return &Etcd{keysAPI}, nil
}

func (c *Etcd) Watch(key string, callback EtcdWatchFunc) error {
	w := c.keysAPI.Watcher(key, &client.WatcherOptions{Recursive: true})

	for {
		resp, err := w.Next(context.TODO())

		if err != nil {
			return errors.Wrap(err, "Failed to iterate.")
		}

		if resp.Node.Dir {
			continue
		}

		callback(resp)
	}
}
