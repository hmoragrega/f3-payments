package payments_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/jsonapi"

	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/hmoragrega/f3-payments/cmd/api/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

var testClient *baloo.Client

var jsonApiContentTypePattern = `application/vnd\.api\+json`

func client() *baloo.Client {
	if testClient == nil {
		testClient = baloo.New(config.NewConfig().GetAPIEndpoint())
	}

	return testClient
}

func reloadFixtures(t *testing.T) {
	r := bytes.NewReader([]byte(getFixtures()))
	l, err := jsonapi.UnmarshalManyPayload(r, reflect.TypeOf(&payment.Payment{}))
	if err != nil {
		t.Fatal(err)
	}

	d, err := config.NewDIC(config.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	if err := d.GetPaymentRepository().DeleteAll(); err != nil {
		t.Fatal(err)
	}

	for _, p := range l {
		if err := d.GetPaymentRepository().Persist(p); err != nil {
			t.Fatal(err)
		}
	}

	//t.Fatal(err)
}

func getErrorResponse(code int, detail string) string {
	return fmt.Sprintf(`{
		"errors": [
			{
				"status": %d,
				"detail": "%s"
			}
		]
	}`, code, detail)
}

func getFixtures() string {
	return `
	{
		"data": [
			{
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
			},
			{
				"type": "payments",
				"id": "502758ff-505f-4d81-b9d2-83aa9c01ebe2",
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
					"processing_time": 1543335309,
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
					"self": "/payments/502758ff-505f-4d81-b9d2-83aa9c01ebe2"
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
	}`
}
