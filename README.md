# Dockerexporter

## Introdução

Esta aplicação tem como objetivo gerar metricas para os indicadores usados em painéis.

## Docker Hub

https://cloud.docker.com/u/jarzamendia/repository/docker/jarzamendia/dockerexporter

## Modo de uso

$ docker exec -it -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION='1.39' jarzamendia/dockerexporter -metric services

Ele irá gerar o resultado em CSV.

## Metricas

### Services

### System

### Nodes

### Containers
