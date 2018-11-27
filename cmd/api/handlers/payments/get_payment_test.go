// +build functional

package payments_test

import (
	"net/http"
	"testing"
)

func TestGetOne(t *testing.T) {
	reloadFixtures(t)
	client().Get("/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43").
		Expect(t).
		Status(http.StatusOK).
		JSON(getPaymentFromFixtures()).
		Done()
}

func TestGetOneNotFoundError(t *testing.T) {
	reloadFixtures(t)
	client().Get("/payments/foo").
		Expect(t).
		Status(http.StatusNotFound).
		JSON(getErrorResponse(http.StatusNotFound, "code=404, message=The payment has not been found")).
		Done()
}

func getPaymentFromFixtures() string {
	return `
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
	}`
}
