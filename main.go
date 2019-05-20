package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jarzamendia/dockerexporter/nodes"
	"github.com/jarzamendia/dockerexporter/services"
	"github.com/jlaffaye/ftp"
)

func main() {

	if len(os.Args) > 5 {
		//TODO - Testar argumentos...

	} else {
		println("Os seguintes parametros devem ser definidos:")
		println("Servidor")
		println("Porta")
		println("Usuario")
		println("Senha")
		println("Pasta")
		println("Ex.: .../dockerexporter ftp.local 21 user@domain senha123 backupfolder")

		os.Exit(0)

	}

	//Aqui recebemos o argumento enviado no Docker run.
	ftpserver := os.Args[1] // Servidor
	ftpport := os.Args[2]   // porta
	user := os.Args[3]      // Usuário
	password := os.Args[4]  // Senha
	folder := os.Args[5]    // Pasta

	fmt.Println("Recebendo informações.")
	var nodeinfo = nodes.Get()
	var serviceinfo = services.Get()

	fmt.Println("Conectando-se ao FTP server.")
	client, err := ftp.Dial(ftpserver + ":" + ftpport)
	if err != nil {
		println("Falha na conexão FTP.")
		panic(err)
	}

	fmt.Println("Autenticando...")
	if err := client.Login(user, password); err != nil {
		println("Falha na autenticação;")
		panic(err)
	}

	fmt.Println("Ok... Acessando pasta.")
	err = client.ChangeDir(folder)
	if err != nil {
		println("Falha ao acessar a pasta")
		panic(err)
	}

	fmt.Println("Criando arquivos")
	serviceByte := strings.Join(serviceinfo, "\r\n")
	nodeByte := strings.Join(nodeinfo, "\r\n")

	sdata := bytes.NewBufferString(serviceByte)
	ndata := bytes.NewBufferString(nodeByte)

	fmt.Println("Salvando arquivos...")
	err = client.Stor("serviceinfo.csv", sdata)
	if err != nil {
		println("Falha ao salvar o arquivo serviceinfo.csv")
		panic(err)
	}

	err = client.Stor("nodeinfo.csv", ndata)
	if err != nil {
		println("Falha ao salvar o arquivo nodeinfo.csv")
		panic(err)
	}
	fmt.Println("Feito.")

}
