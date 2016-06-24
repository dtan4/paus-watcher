package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/coreos/etcd/client"
	"github.com/dtan4/paus-watcher/config"
)

const (
	targetKey = "/paus/users"
)

func callback(resp *client.Response) {
	fmt.Printf("[%s] %s\n", resp.Action, resp.Node.Key)
}

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	etcd, err := NewEtcd(config.EtcdEndpoint)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)

	go func() {
		<-sigch
		os.Exit(0)
	}()

	etcd.Watch(targetKey, callback)
}
