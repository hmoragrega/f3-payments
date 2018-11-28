// +build functional

package payments_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/jsonapi"
)

func TestReplaceCreating(t *testing.T) {
	reloadFixtures(t)
	client().Put("/payments/b27dbd35-7e5a-44d3-81ad-2dda7ccb5a21").
		SetHeader("Content-Type", jsonapi.MediaType).
		JSON(getReplacePayload()).
		Expect(t).
		Type(jsonApiContentTypePattern).
		Status(http.StatusOK).
		JSON(getReplaceCreatingPaymentResponse("b27dbd35-7e5a-44d3-81ad-2dda7ccb5a21")).
		Done()
}

func TestReplaceUpdating(t *testing.T) {
	reloadFixtures(t)
	client().Put("/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43").
		SetHeader("Content-Type", jsonapi.MediaType).
		JSON(getReplacePayload()).
		Expect(t).
		Type(jsonApiContentTypePattern).
		Status(http.StatusOK).
		JSON(getReplaceCreatingPaymentResponse("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")).
		Done()
}

func getReplacePayload() string {
	return `
	{
		"data": {
			"type": "payments",
			"attributes": {
				"amount": 800.11,
				"currency": "USD",
				"beneficiary_party": {
					"account_name": "W Owens",
					"account_number": "31926819",
					"account_number_code": "BBAN",
					"account_type": 0,
					"address": "1 The Beneficiary Localtown SE2",
					"bank_id": "403000",
					"bank_id_code": "GBDSC",
					"name": "Wilfred Jeremiah Owens"
				},
				"debtor_party": {
					"account_name": "EJ Brown Black",
					"account_number": "GB29XABC10161234567801",
					"account_number_code": "IBAN",
					"address": "10 Debtor Crescent Sourcetown NE1",
					"bank_id": "203301",
					"bank_id_code": "GBDSC",
					"name": "Emelia Jane Brown"
				},
				"charges_information": {
					"bearer_code": "SHAR",
					"sender_charges": [{
						"amount": 5.00,
						"currency": "GBP"
					}, {
						"amount": 10.00,
						"currency": "USD"
					}],
					"receiver_charge": {
						"amount": 1.00,
						"currency": "USD"
					}
				},
				"fx": {
					"contract_reference": "FX123",
					"exchange_rate": 2.00000,
					"original_amount": {
						"amount": 200.42,
						"currency": "USD"
					}
				},
				"end_to_end_reference": "Wil piano Jan",
				"numeric_reference": "1002001",
				"payment_id": "123456789012345678",
				"payment_purpose": "Paying for goods/services",
				"payment_scheme": "FPS",
				"payment_type": "Credit",
				"processing_time": 1543110881,
				"reference": "Payment for Em's piano lessons",
				"scheme_payment_sub_type": "InternetBanking",
				"scheme_payment_type": "ImmediatePayment",
				"sponsor_party": {
					"account_number": "56781234",
					"bank_id": "123123",
					"bank_id_code": "GBDSC"
				}
			}
		}
	}`
}

func getReplaceCreatingPaymentResponse(ID string) string {
	return fmt.Sprintf(`
	{
		"data": {
			"type": "payments",
			"id": "%s",
			"attributes": {
				"amount": 800.11,
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
				"processing_time": 1543110881,
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
				"self": "/payments/%s"
			},
			"meta": {
				"organization_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
				"version": "1.0"
			}
		}
	}`, ID, ID)
}
