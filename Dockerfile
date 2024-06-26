FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/.bin/app .bin/app
COPY --from=builder /app/configs configs/

ENV DOCKERIZE_VERSION v0.7.0

RUN apk update --no-cache \
    && apk add --no-cache wget openssl \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apk del wget

EXPOSE 9090
#
#CMD ["./.bin/app"]