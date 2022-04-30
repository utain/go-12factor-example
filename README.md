# GO Example (Web Service)

Trying to implement follow [The Twelve Factor App](https://12factor.net/)

## Dependencies

1. Command-line interface: github.com/spf13/cobra
2. Configuration: github.com/spf13/viper
3. Testing: github.com/stretchr/testify
4. Mocking DB: github.com/DATA-DOG/go-sqlmock
5. ORM: gorm.io/gorm
6. Logging: github.com/op/go-logging
7. HTTP Server: github.com/gin-gonic/gin
8. API Document: github.com/swaggo/swag/cmd/swag

## Project structure

```sh
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── othercmd # example other command line app
│   └── server   # start reading code from here
├── internal
│   ├── api/v1
│   ├── config
|   |-- dto
│   ├── entities
|   |-- errors
|   |-- log
│   ├── services
│   └── utils
├── config
│   └── default.yaml
├── docs
├── dist
│   ├── drawin
│   ├── linux
│   └── windows
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Get started

### Cross platform build environment setup

**macOS**

```sh
# install dep to build binary for linux and windows
brew install FiloSottile/musl-cross/musl-cross
brew install mingw-w64
```

**Command Line**

Run project with docker compose

```sh
docker compose -f dev.yml up --build
```

Run project without build

```sh
go run ./cmd/server [command] --[flag-name]=[flag-value]
```

Generate API Document

```sh
make doc
# open url http://localhost:5000/doc/index.html
```

Build using `make` command

```sh
# Build single binary with specify os
make build[-mac|win|linux]
# Build all os
make all
# Running test
make test
# Start server without build binary file
make run
```

Build with docker

```sh
docker compose build # build docker image
docker compose up # run on docker
# or
docker compose up --build # build and run
docker push [image-name] # public docker image to registry
```


## Configuration

[Viper](https://github.com/spf13/viper#why-viper) uses the following precedence order. Each item takes precedence over the item below it:

- explicit call to Set
- flag
- env
- config
- key/value store
- default

## Example List
- Simple in [main branch](https://github.com/utain/go-12factor-example)
- Port/Adapter in [hexagonal branch](https://github.com/utain/go-12factor-example/tree/hexagonal)
