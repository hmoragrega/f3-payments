FROM golang:1.10.5-alpine3.8

WORKDIR /go/src/github.com/hmoragrega/f3-payments
COPY . .

RUN apk update && \
    apk add --no-cache curl make && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure
RUN make build-api-test

# Create a image with only the app binary
FROM alpine:latest  

WORKDIR /user/local/bin
COPY --from=0 /go/src/github.com/hmoragrega/f3-payments/bin/f3-payments-api .
RUN chmod +x f3-payments-api

CMD ["/user/local/bin/f3-payments-api", "-systemTest", "-test.coverprofile", "/coverage/coverage.functional.cov"] 