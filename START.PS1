$ docker run -it -v \\.\pipe\docker_engine:\\.\pipe\docker_engine -e DOCKER_API_VERSION='1.39'
$ docker run -it -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION='1.39'