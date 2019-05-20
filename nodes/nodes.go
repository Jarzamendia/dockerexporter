package nodes

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

//Get nodes
func Get() []string {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	var list []string

	//Agora corremos pela lista de serviços, adicionando os dados ao slice.
	for _, node := range nodes {

		var hostname = node.Description.Hostname
		var createdAt = node.Meta.CreatedAt
		var plataform = node.Description.Platform.OS
		var engine = node.Description.Engine.EngineVersion

		var role = ""
		var status = ""

		if node.Spec.Role == "manager" {

			role = "Manager"

		} else {

			role = "Worker"
		}

		if node.Status.State == "unknown" {
			status = "Unknown"
		}

		if node.Status.State == "down" {
			status = "Down"
		}

		if node.Status.State == "ready" {
			status = "Ready"
		}

		if node.Status.State == "disconnected" {
			status = "Disconnected"
		}

		var line = hostname + ", " + role + ", " + plataform + ", " + engine + ", " + status + ", " + createdAt.String()

		list = append(list, line)
	}

	return list

}
