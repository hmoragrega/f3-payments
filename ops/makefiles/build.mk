.PHONY: all

API_MAIN := cmd/api/main.go cmd/api/f3api.go
API_EXECUTABLE := bin/f3-payments-api
FLAGS :=-v -race -timeout 10s

test:
	@go test ${FLAGS} ./...

test-unit:
	@go test ${FLAGS} -tags unit ./...

test-functional: 
	@go test ${FLAGS} -tags functional ./...

run:
	@go run ${API_MAIN}

build:
	@go build -o ${API_EXECUTABLE} ${API_MAIN} 