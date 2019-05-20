package services

import (

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"strings"

)

func Get() []string {

	//Primeiro, vamos criar um cliente. Como não definimos configurações, ele usuará o docker.sock.
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	
	//Agora criamos a lista de serviços.
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	//Vamos criar um slice para receber os dados.
	var list []string

	//Agora corremos pela lista de serviços, adicionando os dados ao slice.
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
	
	//Retornamos o slice, terminado por aqui.
	return list

}
