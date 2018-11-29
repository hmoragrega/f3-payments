# F3 API

## Specification
The API is based on the [{json:api}](https://jsonapi.org/) specification using RESTful principles

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
The payment resource can list, create, update and delete payment transactions

**Example:**
```
{
  "data": {
    "type": "payments",
    "id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
    "attributes": {
      "amount": 100.21,
      "beneficiary_party": {
        "name": "Wilfred Jeremiah Owens",
        "address": "1 The Beneficiary Localtown SE2",
        "bank_id": "403000",
        "bank_id_code": "GBDSC",
        "account_name": "W Owens",
        "account_number": "31926819",
        "account_number_code": "BBAN"
      },
      "charges_information": {
        "bearer_code": "SHAR",
        "sender_charges": [{
            "amount": 5,
            "currency": "GBP"
          },
          {
            "amount": 10,
            "currency": "USD"
          }
        ],
        "receiver_charge": {
          "amount": 1,
          "currency": "USD"
        }
      },
      "currency": "USD",
      "debtor_party": {
        "name": "Emelia Jane Brown",
        "address": "10 Debtor Crescent Sourcetown NE1",
        "bank_id": "203301",
        "bank_id_code": "GBDSC",
        "account_name": "EJ Brown Black",
        "account_number": "GB29XABC10161234567801",
        "account_number_code": "IBAN"
      },
      "end_to_end_reference": "Wil piano Jan",
      "fx": {
        "contract_reference": "FX123",
        "exchange_rate": 2,
        "original_amount": {
          "amount": 200.42,
          "currency": "USD"
        }
      },
      "numeric_reference": "1002001",
      "payment_id": "123456789012345678",
      "payment_purpose": "Paying for goods/services",
      "payment_scheme": "FPS",
      "payment_type": "Credit",
      "processing_time": 1542727685,
      "reference": "Payment for Em's piano lessons",
      "scheme_payment_sub_type": "InternetBanking",
      "scheme_payment_type": "ImmediatePayment",
      "sponsor_party": {
        "bank_id": "123123",
        "bank_id_code": "GBDSC",
        "account_number": "56781234"
      }
    },
    "links": {
      "self": "/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
    },
    "meta": {
      "organization_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
      "version": "1.0"
    }
  }
}
```

**Available routes**
- `GET  /payments`
- `GET  /payments/:id`
- `POST  /payments`
- `PUT  /payments/:id`
- `PATCH  /payments/:id`
- `DELETE  /payments/:id`

#### GET  /payments
Lists all the payments

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#bf692430-70e4-4db2-9096-e0c5e579723b)

**Response status**
- `200 OK`: Returns the list of payments
- `422 Unprocessable Entity`: There is a problem with the selected payments
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### GET  /payments/:id
Gets a single payment by the id

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#0d0aae99-38fd-447f-bdef-79e16c962046)

**Response status**
- `200 OK`: Returns the list of payments
- `404 Not Found`: The payment does not exists
- `422 Unprocessable Entity`: There is a problem with the selected payment
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### POST  /payments
Creates a new payment

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#9ce29bef-e2fe-45ec-90d7-44bb7f1970fc)

**Response status**
- `201 Created`: Creates the payment and returns it with the new id
- `400 Bad Request`: The given payment is not valid
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### PUT  /payments/:id
Replaces a payment, or creates a new one if there is no payment for the given id

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#db923beb-6f08-4b2b-a1d0-26fcdd8f1184)

**Response status**
- `200 OK`: Creates or updates the payment
- `400 Bad Request`: The given payment is not valid
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment

#### PATCH  /payments/:id
Updates some attributtes of a payment resource

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#11d1acb3-8b14-4d0e-9879-6a9783af53d6)

**Response status**
- `201 Created`: Creates or updates the payment
- `400 Bad Request`: If the resulting updated payment is not valid
- `404 Not Found`: The payment to update does not exists
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment


#### DELETE  /payments/:id
Deletes a payment

[See examples here](https://documenter.getpostman.com/view/5280062/RzfarrRk#e86a3b1f-283c-4ad3-9562-b1337580595e)

**Response status**
- `204 No Content`: The payment has been deleted
- `500 Internal Server Error`: An unexpected error has happened
- `503 Service Unavailable`: The service is unable to process the request at the moment