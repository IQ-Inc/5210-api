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

With a valid installation of Go, clone the repository in your `$GOPATH/src`
directory.

Quickstart:

```
go run main.go
```

To build the binary, execute `go build` in the directory.

### Tools: `listandscan.go`

`listandscan.go` can be configured to list every DynamoDB table, and
scan for all items in each table. Utility has not been tested against large or
many tables. Use cautiously.

```
go run utils/listandscan.go
```

### Tools: Local DynamoDB

Source the environment variables in `fakeaws.sh`:

```
source fakeaws.sh
```

Process may vary on Windows (not tested). You may be able to change `export` to
`set`, and run it in Powershell / your favorite console emulator. See process 
[here](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
for more information. Specifically, consult the "Environment Variables" section.

Use Docker to create and run a local DynamoDB instance:

```
docker run -d -p [ANOTHER_FAVORITE_PORT]:8000 dwmkerr/dynamodb
```

Once DynamoDB is running, access the shell at `http://localhost:8000/shell`.
Hit the gear ("settings") icon in the menu bar, and set the Access key
ID to the key in `fakeaws.sh`. Create a table by following the tutorial, or by
using the "CreateTable" Javascript template.
