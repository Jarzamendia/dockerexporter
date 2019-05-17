package nodes

import (

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Get() int {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	var nodecount = len(nodes)

	return nodecount

}