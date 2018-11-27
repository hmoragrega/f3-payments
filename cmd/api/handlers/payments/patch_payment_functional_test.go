// +build functional !unit

package payments_test

import (
	"net/http"
	"testing"

	"github.com/google/jsonapi"
)

func TestPatch(t *testing.T) {
	reloadFixtures(t)
	client().Patch("/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43").
		SetHeader("Content-Type", jsonapi.MediaType).
		JSON(getPatchPaymentPayload()).
		Expect(t).
		Type(jsonApiContentTypePattern).
		Status(http.StatusOK).
		JSON(getPatchPaymentResponse()).
		Done()
}

func getPatchPaymentPayload() string {
	return `
	{
		"data": {
			"type": "payments",
			"attributes": {
				"amount": 500.32,
				"beneficiary_party": {
					"account_name": "foo"
				}
			}
		}
	}`
}

func getPatchPaymentResponse() string {
	return `
	{
		"data": {
			"type": "payments",
			"id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			"attributes": {
				"amount": 500.32,
				"beneficiary_party": {
					"name": "Wilfred Jeremiah Owens",
					"address": "1 The Beneficiary Localtown SE2",
					"bank_id": "403000",
					"bank_id_code": "GBDSC",
					"account_name": "foo",
					"account_number": "31926819",
					"account_number_code": "BBAN"
				},
				"charges_information": {
					"bearer_code": "SHAR",
					"sender_charges": [
						{
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
	}`
}
