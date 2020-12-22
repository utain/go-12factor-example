FROM golang:alpine
RUN apk add git gcc make
RUN mkdir /app
WORKDIR /app
# RUN go get -u honnef.co/go/tools
ADD go.mod .
ADD go.sum .
RUN go mod download

ENV GIN_MODE=debug

CMD sh -c "go run ./cmd/server start"