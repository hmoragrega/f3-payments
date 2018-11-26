FROM golang:1.10-alpine3.8

RUN apk update && \
    apk add --no-cache mongodb-tools make docker curl netcat-openbsd py-pip && \
    pip install 'docker-compose==1.8.0' && \
    curl -O -L https://github.com/ufoscout/docker-compose-wait/releases/download/2.4.0/wait && \
    mv wait /usr/local/bin