.PHONY: all

API_MAIN := server/api.go
API_EXECUTABLE := bin/f3-payments-api
FLAGS := -race

test:
	@go test -v ${FLAGS} ./...

test-unit: 
	@go test -v -tags unit ${FLAGS} ./...

test-functional: 
	@go test -v -tags functional ${FLAGS} ./...

run:
	@go run ${FLAGS} ${API_MAIN}

build:
	@go build -o ${API_EXECUTABLE} ${API_MAIN} 