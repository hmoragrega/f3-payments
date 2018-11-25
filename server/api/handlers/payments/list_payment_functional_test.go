package payments_test

import (
	"net/http"
	"testing"
)

func TestGetList(t *testing.T) {
	client().Get("/payments").
		Expect(t).
		Status(http.StatusOK).
		JSON(getPaymentListFromFixtures()).
		Done()
}

func getPaymentListFromFixtures() string {
	return `
	{
		"data": [
			{
				"type": "payments",
				"id": "502758ff-505f-4d81-b9d2-83aa9c01ebe2",
				"attributes": {
					"amount": 100.21,
					"beneficiary_party": {
						"Name": "Wilfred Jeremiah Owens",
						"Address": "1 The Beneficiary Localtown SE2",
						"BankID": "403000",
						"BankIDCode": "GBDSC",
						"AccountName": "W Owens",
						"AccountNumber": "31926819",
						"AccountNumberCode": "BBAN",
						"AccountType": 0
					},
					"charges_information": {
						"BearerCode": "SHAR",
						"SenderCharges": [
							{
								"Amount": 5,
								"Currency": "GBP"
							},
							{
								"Amount": 10,
								"Currency": "USD"
							}
						],
						"ReceiverCharge": {
							"Amount": 1,
							"Currency": "USD"
						}
					},
					"currency": "USD",
					"debtor_party": {
						"Name": "Emelia Jane Brown",
						"Address": "10 Debtor Crescent Sourcetown NE1",
						"BankID": "203301",
						"BankIDCode": "GBDSC",
						"AccountName": "EJ Brown Black",
						"AccountNumber": "GB29XABC10161234567801",
						"AccountNumberCode": "IBAN",
						"AccountType": 0
					},
					"end_to_end_reference": "Wil piano Jan",
					"fx": {
						"ContractReference": "FX123",
						"ExchangeRate": 2,
						"OriginalAmount": {
							"Amount": 200.42,
							"Currency": "USD"
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
						"Name": "",
						"Address": "",
						"BankID": "123123",
						"BankIDCode": "GBDSC",
						"AccountName": "",
						"AccountNumber": "56781234",
						"AccountNumberCode": "",
						"AccountType": 0
					}
				},
				"links": {
					"self": "/payments/502758ff-505f-4d81-b9d2-83aa9c01ebe2"
				},
				"meta": {
					"organization_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
					"version": "1.0"
				}
			},
			{
				"type": "payments",
				"id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
				"attributes": {
					"amount": 100.21,
					"beneficiary_party": {
						"Name": "Wilfred Jeremiah Owens",
						"Address": "1 The Beneficiary Localtown SE2",
						"BankID": "403000",
						"BankIDCode": "GBDSC",
						"AccountName": "W Owens",
						"AccountNumber": "31926819",
						"AccountNumberCode": "BBAN",
						"AccountType": 0
					},
					"charges_information": {
						"BearerCode": "SHAR",
						"SenderCharges": [
							{
								"Amount": 5,
								"Currency": "GBP"
							},
							{
								"Amount": 10,
								"Currency": "USD"
							}
						],
						"ReceiverCharge": {
							"Amount": 1,
							"Currency": "USD"
						}
					},
					"currency": "USD",
					"debtor_party": {
						"Name": "Emelia Jane Brown",
						"Address": "10 Debtor Crescent Sourcetown NE1",
						"BankID": "203301",
						"BankIDCode": "GBDSC",
						"AccountName": "EJ Brown Black",
						"AccountNumber": "GB29XABC10161234567801",
						"AccountNumberCode": "IBAN",
						"AccountType": 0
					},
					"end_to_end_reference": "Wil piano Jan",
					"fx": {
						"ContractReference": "FX123",
						"ExchangeRate": 2,
						"OriginalAmount": {
							"Amount": 200.42,
							"Currency": "USD"
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
						"Name": "",
						"Address": "",
						"BankID": "123123",
						"BankIDCode": "GBDSC",
						"AccountName": "",
						"AccountNumber": "56781234",
						"AccountNumberCode": "",
						"AccountType": 0
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
				"id": "09fe827a-b3c2-4437-b999-6c0e780c0983",
				"attributes": {
					"amount": 100.21,
					"beneficiary_party": {
						"Name": "Wilfred Jeremiah Owens",
						"Address": "1 The Beneficiary Localtown SE2",
						"BankID": "403000",
						"BankIDCode": "GBDSC",
						"AccountName": "W Owens",
						"AccountNumber": "31926819",
						"AccountNumberCode": "BBAN",
						"AccountType": 0
					},
					"charges_information": {
						"BearerCode": "SHAR",
						"SenderCharges": [
							{
								"Amount": 5,
								"Currency": "GBP"
							},
							{
								"Amount": 10,
								"Currency": "USD"
							}
						],
						"ReceiverCharge": {
							"Amount": 1,
							"Currency": "USD"
						}
					},
					"currency": "USD",
					"debtor_party": {
						"Name": "Emelia Jane Brown",
						"Address": "10 Debtor Crescent Sourcetown NE1",
						"BankID": "203301",
						"BankIDCode": "GBDSC",
						"AccountName": "EJ Brown Black",
						"AccountNumber": "GB29XABC10161234567801",
						"AccountNumberCode": "IBAN",
						"AccountType": 0
					},
					"end_to_end_reference": "Wil piano Jan",
					"fx": {
						"ContractReference": "FX123",
						"ExchangeRate": 2,
						"OriginalAmount": {
							"Amount": 200.42,
							"Currency": "USD"
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
						"Name": "",
						"Address": "",
						"BankID": "123123",
						"BankIDCode": "GBDSC",
						"AccountName": "",
						"AccountNumber": "56781234",
						"AccountNumberCode": "",
						"AccountType": 0
					}
				},
				"links": {
					"self": "/payments/09fe827a-b3c2-4437-b999-6c0e780c0983"
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
