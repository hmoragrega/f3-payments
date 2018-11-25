package payments_test

import (
	"net/http"
	"testing"

	"github.com/hmoragrega/f3-payments/server/api"
	baloo "gopkg.in/h2non/baloo.v3"
)

var client = getTestClient()

func TestGetCollection(t *testing.T) {
	client.Get("/payments").
		Expect(t).
		Status(http.StatusOK).
		JSON(getDummyCollection(t)).
		Done()
}

func getTestClient() *baloo.Client {
	return baloo.New(api.NewF3API().GetAddress())
}

func getDummyCollection(t *testing.T) []byte {
	return []byte(`
	{
		"data": [
			{
				"type": "payments",
				"id": "foo",
				"attributes": {
					"amount": 0,
					"beneficiary_party": {
						"Name": "",
						"Address": "",
						"BankID": "",
						"BankIDCode": "",
						"AccountName": "",
						"AccountNumber": "",
						"AccountNumberCode": "",
						"AccountType": 0
					},
					"charges_information": {
						"BearerCode": "",
						"SenderCharges": null,
						"ReceiverCharge": {
							"Amount": 0,
							"Currency": ""
						}
					},
					"currency": "",
					"debtor_party": {
						"Name": "",
						"Address": "",
						"BankID": "",
						"BankIDCode": "",
						"AccountName": "",
						"AccountNumber": "",
						"AccountNumberCode": "",
						"AccountType": 0
					},
					"end_to_end_reference": "",
					"fx": {
						"ContractReference": "",
						"ExchangeRate": 0,
						"OriginalAmount": {
							"Amount": 0,
							"Currency": ""
						}
					},
					"numeric_reference": "",
					"payment_id": "",
					"payment_purpose": "",
					"payment_scheme": "",
					"payment_type": "",
					"reference": "",
					"scheme_payment_sub_type": "",
					"scheme_payment_type": "",
					"sponsor_party": {
						"Name": "",
						"Address": "",
						"BankID": "",
						"BankIDCode": "",
						"AccountName": "",
						"AccountNumber": "",
						"AccountNumberCode": "",
						"AccountType": 0
					}
				},
				"links": {
					"self": "/payments/foo"
				},
				"meta": {
					"organization_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
					"version": "1.0"
				}
			}
		],
		"links": {
			"self": "/payments"
		}
	}
	`)
}
