# GO Example (Web Service)

Trying to implement follow [The Twelve Factor App](https://12factor.net/)

### Dependencies
1. Command-line interface: github.com/spf13/cobra
2. Configuration: github.com/spf13/viper
3. Testing: github.com/stretchr/testify
4. Mocking DB: github.com/DATA-DOG/go-sqlmock
5. ORM: github.com/jinzhu/gorm
6. Logging: github.com/sirupsen/logrus [todo]
7. HTTP Server: github.com/gin-gonic/gin

### Project structure
```
.
├── Dockerfile
├── Makefile
├── README.md
├── app
│   ├── api/v1
│   ├── config
│   ├── log
│   ├── models
│   ├── services
│   └── utils
├── cmd
│   ├── root.go
│   ├── start.go
│   └── version.go
├── config
│   └── default.yaml
├── docker-compose.yml
├── exam.db
├── go.mod
├── go.sum
└── main.go
```

### Get started

**Cross platform build environment setup**

macOS
```sh
# install dep to build binary for linux and windows
brew install FiloSottile/musl-cross/musl-cross
brew install mingw-w64
```

Linux/Windows
```sh
# todo
```


**Command**

Run project without build
```sh
go run . [command] --[flag-name]=[flag-value]
```

Build using `make` command
```sh
# Build single binary with specify os
make build[-mac|win|linux]
# Build all os
make all
# Running test
make test
```

Build with docker
```sh
docker-compose build # build docker image
docker-compose up # run on docker
# or 
docker-compose up --build # build and run
docker push [image-name]
```


### Configuration
[Viper](https://github.com/spf13/viper#why-viper) uses the following precedence order. Each item takes precedence over the item below it:

- explicit call to Set
- flag
- env
- config
- key/value store
- default