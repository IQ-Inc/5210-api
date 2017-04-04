# 5210-api

Backend for 5210 Applicaiton

## Usage: Docker

From this directory, with a valid installation of Docker, run

```
docker build -t letsmove .
```

Docker compiles the go code inside the container and creates a custom image
devoted to running your application. From there, you may create a container:

```
docker run -it -p [YOUR_FAVORITE_PORT]:9000 --name letsmove-api letsmove
```

Use CTRL-C to exit the running container. Use `docker start -a letsmove-api` 
to start and re-attach to the container.

## Usage: Go

With a valid installation of Go, clone the repository in your `$GOPATH/src` directory.

Quickstart:

```
go run main.go
```

To build the binary, execute `go build` in the directory.