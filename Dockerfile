FROM golang:alpine as gobuilder

ENV GIN_MODE=release

RUN apk add git gcc make

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN make

FROM alpine:latest as ginserv
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/

# Default server port
EXPOSE 8000

COPY --from=gobuilder /app/dist/server /bin/server

CMD server gin

FROM alpine:latest as fiberserv
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/

# Default server port
EXPOSE 8000

COPY --from=gobuilder /app/dist/server /bin/server

CMD server fiber