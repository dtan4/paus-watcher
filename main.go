package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/coreos/etcd/client"
)

func callback(resp *client.Response) {
	fmt.Printf("[%s] %s\n", resp.Action, resp.Node.Key)
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	endpoint := os.Args[1]
	etcd, err := NewEtcd(endpoint)

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

	etcd.Watch("/paus/users/", callback)
}
