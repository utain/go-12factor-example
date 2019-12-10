# Go parameters
export CGO_ENABLED:=0
export GOCMD:=go
export GOARCH:=amd64
export BINARY_NAME:=go-example
export BINARY_NAME_LINUX:=./dist/linux/$(BINARY_NAME)
export BINARY_NAME_MACOS:=./dist/drawin/$(BINARY_NAME)
export BINARY_NAME_WIN:=./dist/windows/$(BINARY_NAME).exe
all: test build
build: 
	GOOS=linux $(GOCMD) build -o $(BINARY_NAME_LINUX)
	GOOS=darwin $(GOCMD) build -o $(BINARY_NAME_MACOS)
	GOOS=windows $(GOCMD) build -o $(BINARY_NAME_WIN)
test: 
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
	rm -rf ./dist
download:
	go mod download
