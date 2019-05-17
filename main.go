package main

import (
	"fmt"
	"os"

	"github.com/jarzamendia/dockerexporter/containers"
	"github.com/jarzamendia/dockerexporter/system"
	"github.com/jarzamendia/dockerexporter/services"
	"github.com/jarzamendia/dockerexporter/nodes"

)

func main() {

	switch os.Args[1] {

	case "services":
		var serviceinfo = services.Get()
		for _, service := range serviceinfo {

			fmt.Println(service)

		}

	case "nodes":
		var nodeinfo = nodes.Get()
		fmt.Println(nodeinfo)

	case "containers":
		var containerinfo = containers.Get()
		fmt.Println(containerinfo)

	case "system":
		var systeminfo = system.Get()
		fmt.Println(systeminfo)

	default:
		fmt.Println("As opções possiveis são:")
		fmt.Println("services")
		fmt.Println("nodes")
		fmt.Println("containers")
		fmt.Println("system")
	}
	
}
