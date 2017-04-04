package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	DOCKER_NETWORK = "bfscalp"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}
	if !findNetwork(cli, DOCKER_NETWORK) {
		createNetwork(DOCKER_NETWORK, cli)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.Names, container.Image)
	}
}

func findNetwork(network string, networkList []types.NetworkResource) {
	for _, network := range networks {
		if network.Name == DOCKER_NETWORK {

		}
	}
}
