FROM golang:alpine as gobuilder

ENV GIN_MODE=release

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./server ./cmd/server

FROM alpine:latest
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/
EXPOSE 5000
COPY --from=gobuilder /app/server /bin/server
ADD ./wait-for ./wait-for
ENV DATABASE_TYPE="postgres"
# ENV DATABASE_URL="host=db port=5432 user=example dbname=example password=P@55w0rd sslmode=disable"
CMD ./wait-for $DATABASE_HOST:$DATABASE_PORT -- echo "postgres is up" && server migrate -d && server start