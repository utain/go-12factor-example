TIME=$(shell date "+%Y.%m.%dT%H:%M:%S")
COMMIT=$(shell git rev-parse --short HEAD)
VERSION=v1.0.0

build:
	go build -ldflags="-X 'github.com/utain/go/example/internal/version.Version=$(VERSION)' -X 'github.com/utain/go/example/internal/version.Time=$(TIME)' -X 'github.com/utain/go/example/internal/version.Commit=$(COMMIT)'" -o dist/server .

.PHONY: build