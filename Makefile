export CGO_ENABLED=1
export GOCMD=go
export GOARCH=amd64
export BINARY_NAME=go-example
export BINARY_NAME_LINUX=./dist/linux/$(BINARY_NAME)
export BINARY_NAME_MACOS=./dist/drawin/$(BINARY_NAME)
export BINARY_NAME_WIN=./dist/windows/$(BINARY_NAME).exe
all: test build
build-linux: 
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" GOOS=linux $(GOCMD) build -i -a -v -o $(BINARY_NAME_LINUX)
build-mac:
	GOOS=darwin $(GOCMD) build -i -a -v -o $(BINARY_NAME_MACOS)
build-win:
	CC="x86_64-w64-mingw32-gcc" GOOS=windows $(GOCMD) build -i -a -v -o $(BINARY_NAME_WIN)
build: build-mac build-win build-linux
test: 
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
	rm -rf ./dist
download:
	go mod download
