FROM golang:1.13-alpine

ENV GIN_MODE=release

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -i -a -v .

FROM alpine:latest
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 5000
COPY --from=0 /app/go-example /bin/go-example
CMD ["/bin/go-example", "start"]