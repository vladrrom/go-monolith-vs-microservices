## Installation

- Checkout source code

`$ git clone https://github.com/vladrrom/go-grpc.git`

- gRPC installation

`$ go get -u google.golang.org/grpc`

- Compilier protoc installation

`$ go get github.com/golang/protobuf@v1.4`

## Commands

- Proto file generate

`$ protoc -I proto proto/piservice.proto --go_out=plugins=grpc:proto/`

## Local project

- Building and start server

```sh
$ go build -race -ldflags "-s -w" -o bin/server server/main.go
$ bin/server
```

- Building and start client

```sh
$ go build -ldflags "-s -w" -o bin/client client/main.go
$ bin/client 500000
```

## Docker container

- Docker run. Building and start server in container

```sh
$ docker-compose up
$ docker-compose logs -f --tail="50" CalcPiGRPC
$ docker images
```

- Building and start client

```sh
$ go build -ldflags "-s -w" -o bin/client client/main.go
$ bin/client 500000
```
