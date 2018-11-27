.PHONY: all

API_MAIN := cmd/api/main.go cmd/api/f3api.go
API_EXECUTABLE := bin/f3-payments-api
FLAGS :=-v -race -timeout 10s
GO_TEST = GOCACHE=off go test
GO_COVER = -covermode=count -coverpkg=./...

test:
	@${GO_TEST} ${FLAGS} -tags='unit functional' ./...

test-unit:
	@${GO_TEST} ${FLAGS} -tags=unit ./...

test-functional:
	@${GO_TEST} ${FLAGS} -tags=functional ./...

coverage:
	@${GO_TEST} -v -timeout 10s -tags='unit functional' ${GO_COVER} -coverprofile=coverage.unit.cov ./...

run:
	@go run ${API_MAIN}

build-api:
	@go build -o ${API_EXECUTABLE} ${API_MAIN}

build-api-test:
	@go test -c -tags=api ${GO_COVER} -o ${API_EXECUTABLE} ./cmd/api