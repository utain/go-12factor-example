CGO_ENABLED=1
GOCMD=go
GOARCH=amd64
BINARY_NAME=go-example
BINARY_NAME_FILE=./dist/$(BINARY_NAME)
BINARY_NAME_LINUX=./dist/linux/$(BINARY_NAME)
BINARY_NAME_MACOS=./dist/drawin/$(BINARY_NAME)
BINARY_NAME_WIN=./dist/windows/$(BINARY_NAME).exe
build:
	$(GOCMD) build -i -a -v -o $(BINARY_NAME_FILE)
build-linux:
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" GOOS=linux $(GOCMD) build -i -a -v -o $(BINARY_NAME_LINUX)
build-mac:
	GOOS=darwin $(GOCMD) build -i -a -v -o $(BINARY_NAME_MACOS)
build-win:
	CC="x86_64-w64-mingw32-gcc" GOOS=windows $(GOCMD) build -i -a -v -o $(BINARY_NAME_WIN)
test:
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
	rm -rf ./dist
download:
	go mod download
build-all: build-mac build-win build-linux
all: test build-all
