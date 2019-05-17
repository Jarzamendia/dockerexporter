package system

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
	
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	var servicecount = len(services)
	
	return servicecount

}
