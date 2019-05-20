package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jarzamendia/dockerexporter/services"
)

func main() {

	//Service
	fmt.Println("Gerando lista de servicos.")
	var serviceinfo = services.Get()

	fmt.Println("Testando arquivo serviceinfo.csv.")
	if _, err := os.Stat("serviceinfo.csv"); os.IsNotExist(err) {

		fmt.Println("Arquivo não encontrado... Criando...")
		file, err := os.Create("serviceinfo.csv")
		checkError("Cannot create file", err)

		defer file.Close()
		fmt.Println("Arquivo criado.")

	}

	fmt.Println("Abrindo arquivo...")
	file, err := os.Open("serviceinfo.csv")
	checkError("Cannot create file", err)
	fmt.Println("Arquivo aberto.")

	fmt.Println("Preparando escrita.")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("Escrevendo services.")
	writer.Write(serviceinfo)
	fmt.Println("Service escrito.")

	time.Sleep(1000 * time.Second)

	//Node
	//var nodeinfo = nodes.Get()

	//for _, node := range nodeinfo {

	//	fmt.Println(node)

	//}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
