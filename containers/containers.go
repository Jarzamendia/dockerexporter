package containers

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
	
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	var containerscount = len(containers)

	return containerscount
	
}