package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/coreos/etcd/client"
	"github.com/dtan4/paus-watcher/config"
	p "github.com/dtan4/paus-watcher/provider"
	"github.com/dtan4/paus-watcher/store"
)

func callback(provider p.Provider, resp *client.Response) error {
	// Action: get, set, delete, update, create, compareAndSwap, compareAndDelete and expire.
	if err := provider.Notify(resp.Action, resp.Node.Key, resp.Node.Value); err != nil {
		return err
	}

	fmt.Printf("[%s] %s\n", resp.Action, resp.Node.Key)

	return nil
}

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	etcd, err := store.NewEtcd(config.EtcdEndpoint)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	provider := p.NewProvider(config)

	if provider == nil {
		fmt.Fprintln(os.Stderr, "No provider can be used.")
		os.Exit(1)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)

	go func() {
		<-sigch
		os.Exit(0)
	}()

	if err := etcd.Watch(config.TargetKey, provider, callback); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
