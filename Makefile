# OS
OSNAME 				    :=
BINARY_NAME_FILE  :=
ifeq ($(OS),Windows_NT)
	OSNAME=windows
else
	UNAME_S :=$(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSNAME=linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSNAME=drawin
	endif
endif

# Env
CGO_ENABLED=1
GOCMD=go
GOARCH=amd64
BINARY_NAME=server
BINARY_NAME_FILE =./dist/$(OSNAME)/
BINARY_NAME_LINUX=./dist/linux/
BINARY_NAME_MACOS=./dist/drawin/
BINARY_NAME_WIN=./dist/windows/
GIT_COMMIT=$(shell git rev-list -1 HEAD)
VERSION=$(shell date "+%Y.%m.%d.%H:%M:%S")
GIT_TAG=$(shell git describe --all)
BUILD_FLAGS=-i -v -ldflags "-X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(GIT_TAG) -X main.BuildDate=$(VERSION)"
prebuild:
	mkdir -p ./dist/$(OSNAME)/
prebuild-all:
	mkdir -p $(BINARY_NAME_LINUX)
	mkdir -p $(BINARY_NAME_MACOS)
	mkdir -p $(BINARY_NAME_WIN)
_dist_os:
	$(GOCMD) build $(BUILD_FLAGS) -o $(BINARY_NAME_FILE) ./cmd/...
build: prebuild _dist_os
build-linux:
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" GOOS=linux $(GOCMD) build $(BUILD_FLAGS) -o $(BINARY_NAME_LINUX) ./cmd/...
build-mac:
	GOOS=darwin $(GOCMD) build $(BUILD_FLAGS) -o $(BINARY_NAME_MACOS) ./cmd/...
build-win:
	CC="x86_64-w64-mingw32-gcc" GOOS=windows $(GOCMD) build $(BUILD_FLAGS) -o $(BINARY_NAME_WIN) ./cmd/...
test:
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean ./...
	rm -rf ./dist/
download:
	go mod download
build-all: build-mac build-win build-linux
all: test prebuild-all build-all
run:
	$(GOCMD) run ./cmd/server start
build-image:
	docker-compose build
run-docker:
	docker-compose up --build