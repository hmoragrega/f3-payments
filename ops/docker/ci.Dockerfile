FROM golang:1.10-alpine3.8

RUN apk update && \
    apk add --no-cache mongodb-tools make docker curl netcat-openbsd py-pip && \
    pip install 'docker-compose==1.8.0'