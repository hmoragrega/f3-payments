FROM golang:1.10-alpine3.8

RUN apk update && \
    apk add --no-cache mongodb-tools make docker curl

RUN curl -Lo /usr/local/bin/docker-compose https://github.com/docker/compose/releases/download/1.8.0/docker-compose-Linux-x86_64 && \
    chmod a+rx /usr/local/bin/docker-compose