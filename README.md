# Dockerexporter

## Introdução

Esta aplicação tem como objetivo gerar metricas para os indicadores usados em painéis da SGI. Ela é desenvolvida em Golang.

Ela acessa a API do Docker, através da SDK do Docker para Golang. O golang foi utilizado por poder ser facilmente compilado para Windows e Linux. Teoricamente basta ajustar a Dockerfile para trocar de sistema operacional.

## Docker Hub

https://cloud.docker.com/u/jarzamendia/repository/docker/jarzamendia/dockerexporter

## Modo de uso

$ docker exec -it -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION='1.39' jarzamendia/dockerexporter services
$ docker exec -it -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION='1.39' jarzamendia/dockerexporter nodes

Ele irá gerar o resultado em CSV.

## Metricas

As seguintes metricas são recebidas:

### Services

Nome, Modo, Imagem, CriadoEm

### Nodes

Hostname, Papel, Plataforma(OS), Versão, Status, CriadoEm