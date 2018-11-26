# F3 API

## Specification
The API is based on the [(json:api}](https://jsonapi.org/) specification using RESTful principles

## Authentication/Authorization
Currently there is none in place to keep things simple

## Errors
The errors follow the next example

```
{
  "errors": [
    {
      "status": 400,
      "detail": "Error message"
    }
  ]
}

```
## Resources
These are the avaliable resource in the API

### Payment
The payment resource can list, create, update and delete transactions

**Available routes**
- `GET /payment`
- `GET /payment/:id`
- `POST /payment`
- `PUT /payment/:id`
- `PATCH /payment/:id`
- `DELETE /payment/:id`

#### GET /payment
Lists all the payments

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#bf692430-70e4-4db2-9096-e0c5e579723b)

**Response status**
- `200 OK`: Returns the list of payments
- `422 Unprocessable Entity`: There is a problem with the selected payments
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### GET /payment/:id
Gets a single payment by the id

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#0d0aae99-38fd-447f-bdef-79e16c962046)

**Response status**
- `200 OK`: Returns the list of payments
- `404 Not Found`: The payment does not exists
- `422 Unprocessable Entity`: There is a problem with the selected payment
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### POST /payment
Creates a new payment

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#9ce29bef-e2fe-45ec-90d7-44bb7f1970fc)

**Response status**
- `201 Created`: Creates the payment and returns it with the new id
- `400 Bad Request`: The given payment is not valid
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### PUT /payment/:id
Replaces a payment, or creates a new one if there is no payment for the given id

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#db923beb-6f08-4b2b-a1d0-26fcdd8f1184)

**Response status**
- `200 OK`: Creates or updates the payment
- `400 Bad Request`: The given payment is not valid
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### PATCH /payment/:id
Updates some attributtes of a payment resource

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#11d1acb3-8b14-4d0e-9879-6a9783af53d6)

**Response status**
- `201 Created`: Creates or updates the payment
- `400 Bad Request`: If the resulting updated payment is not valid
- `404 Not Found`: The payment to update does not exists
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment


#### DELETE /payment/:id
Deletes a payment

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#e86a3b1f-283c-4ad3-9562-b1337580595e)

**Response status**
- `204 No Content`: The payment has been deleted
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment