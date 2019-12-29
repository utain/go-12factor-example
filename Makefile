# OS
OSFLAG 				    :=
BINARY_NAME_FILE  :=
ifeq ($(OS),Windows_NT)
	OSFLAG=windows
else
	UNAME_S :=$(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSFLAG=linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSFLAG=drawin
	endif
endif

hello:
	@echo $(OSFLAG)

# Env
CGO_ENABLED=1
GOCMD=go
GOARCH=amd64
BINARY_NAME=server
BINARY_NAME_FILE =./dist/$(OSFLAG)/$(BINARY_NAME)
BINARY_NAME_LINUX=./dist/linux/$(BINARY_NAME)
BINARY_NAME_MACOS=./dist/drawin/$(BINARY_NAME)
BINARY_NAME_WIN=./dist/windows/$(BINARY_NAME).exe
build:
	$(GOCMD) build -i -v -o $(BINARY_NAME_FILE) ./cmd/...
build-linux:
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" GOOS=linux $(GOCMD) build -i -v -o $(BINARY_NAME_LINUX) ./cmd/...
build-mac:
	GOOS=darwin $(GOCMD) build -i -v -o $(BINARY_NAME_MACOS) ./cmd/...
build-win:
	CC="x86_64-w64-mingw32-gcc" GOOS=windows $(GOCMD) build -i -v -o $(BINARY_NAME_WIN) ./cmd/...
test:
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
	rm -rf ./dist
download:
	go mod download
build-all: build-mac build-win build-linux
all: test build-all
