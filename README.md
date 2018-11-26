# F3 payments
[![Go Report Card](https://goreportcard.com/badge/github.com/hmoragrega/f3-payments)](https://goreportcard.com/report/github.com/hmoragrega/f3-payments)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/hmoragrega/f3-payments/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/hmoragrega/f3-payments/?branch=master)
[![Build Status](https://scrutinizer-ci.com/g/hmoragrega/f3-payments/badges/build.png?b=master)](https://scrutinizer-ci.com/g/hmoragrega/f3-payments/build-status/master)

## Pre-requisites
- [Go v1.10](https://golang.org/)
- [Docker](https://docs.docker.com/install/#releases)
- [Docker-compose](https://docs.docker.com/compose/install/)
- [GNU make](https://www.gnu.org/software/make/)

### Configuration
The API can be configured trough environmental variables
- `F3_API_IP`: The IP address of the to bind server
- `F3_API_PORT`: The IP address of the to binf the API service (default: `80`)
- `F3_API_LOG_FORMAT`: A log format compatible with (logrus)[https://github.com/sirupsen/logrus]
- `F3_API_MONGO_SERVER`: The address of the mongo server (default: `127.0.0.1:27017`)
- `F3_API_MONGO_USER`: The username to use to connect to mongo (default: `root`)
- `F3_API_MONGO_PASS`: The password to use to connect to mongo (default: `demo`)
- `F3_API_MONGO_DATABASE`: The API database (default: `f3api`)
- `F3_API_MONGO_AUTH_DB`: The name of the mongo admin database (default: `admin`)

## Development environment
The development environment can be managed trough make commands

### Dependencies
The API requires this infrastructure
- [MongoDB](https://www.mongodb.com/)

To start the mongo server we will use a docker image managed with docker-compose
```
make dev
```

To stop it
```
make stop
```

To removed it
```
make down
```

### Run
This command will compile and run the API for development
```
make run
```
### Build
This command will compile and build the API server binary and stored it as `bin/f3-payment-api`
```
make build
```

## Test
There's 2 tests suites:
- Unit
- Integration

### Unit 
This command will run the unit tests
```
make test-unit
```

### Functional 
**NOTE** The functional tests require a the API server to be [up & running](#run), the tests will use the same [config environment variables](#configuration) to determine the API address
```
make test-functional
```

## TODO
- Use HTTPS protocol
- Use authentication like a custom JWT based or full-fledge Oauth2
- API versioning
- Pagination
- Filter/Search the payment lists
- Reconnect logic on MongoDB repository