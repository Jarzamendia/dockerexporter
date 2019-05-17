package services

import (

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"strings"

)

func Get() []string {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	//Name,ID,Version,Mode,Image
	var list []string

	for _, service := range services {

		var name = service.Spec.Name

		var mode = ""

		if service.Spec.Mode.Replicated == nil {
			mode = "Global"
		} else {
			mode = "Replicated"
		}

		var image =  strings.Split(service.Spec.TaskTemplate.ContainerSpec.Image, "@")[0]

		var createdAt = service.Meta.CreatedAt
		
		var line = name + ", " + mode + ", " + image + ", " + createdAt.String()
		
		list = append(list, line)
	}
	
	return list

}
