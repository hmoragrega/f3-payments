.PHONY: all

API_MAIN := cmd/api/main.go cmd/api/f3api.go
API_EXECUTABLE := bin/f3-payments-api
FLAGS :=-v -race -timeout 10s
GO_TEST = GOCACHE=off go test

test:
	@${GO_TEST} ${FLAGS} -tags='unit functional' ./...

test-unit:
	@${GO_TEST} ${FLAGS} -tags=unit ./...

test-functional:
	@${GO_TEST} ${FLAGS} -tags=functional ./...

coverage:
	@${GO_TEST} ${FLAGS} -tags='unit functional' -coverpkg=./... -coverprofile=coverage.txt ./... 

run:
	@go run ${API_MAIN}

build-api:
	@go build -o ${API_EXECUTABLE} ${API_MAIN} 