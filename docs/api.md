# F3 API

## Specification
The API is based on the [(json:api}](https://jsonapi.org/) specification using RESTful principles

## Authentication/Authorization
Currently there is none in place to keep things simple

## Resources
These are the avaliable resource in the API

### Payment
The payment resource can list, create, update and delete transactions

See examples at: 

**Available routes**
- `GET /payment`
- `GET /payment/:id`
- `POST /payment`
- `PUT /payment/:id`
- `PATCH /payment/:id`
- `DELETE /payment/:id`

### GET /payment
Lists all the payments

See [examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#bf692430-70e4-4db2-9096-e0c5e579723b)

### Posible response:
- `200 OK` Returns the list of payments
- `422 Unprocessable Entity`: There has been a problem processing the payments data
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment