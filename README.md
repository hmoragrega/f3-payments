# F3 payments

## Pre-requisites
- [Docker](https://docs.docker.com/install/#releases)
- [Docker-compose](https://docs.docker.com/compose/install/)
- [GNU make](https://www.gnu.org/software/make/)

### Configuration
The API can be configured trough environmental variables
- `F3_API_IP`: The IP address of the to bind server
- `F3_API_PORT`: The IP address of the to binf the API service (default: `80`)

## Development environment
The development environment can be managed trough make commands

### Run
This command will run the API for development
```
make run
```
### Build
This command will build the API server binary and stored it as `bin/f3-payment-api`
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