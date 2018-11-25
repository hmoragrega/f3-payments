// +build unit !functional

package payment

import (
	"bytes"
	"testing"
	"time"

	"github.com/Slemgrim/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestJSONApiMarshaling(t *testing.T) {
	expected := new(Payment)
	err := jsonapi.UnmarshalPayload(bytes.NewReader(getJSONPayment()), expected)
	if err != nil {
		t.Fatal("The payment fixture could not be unmarshaled as a json:api struct", err)
	}

	payment := &Payment{
		ID:       "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		Amount:   100.21,
		Currency: "USD",
		Beneficiary: Party{
			Name:              "Wilfred Jeremiah Owens",
			Address:           "1 The Beneficiary Localtown SE2",
			BankID:            "403000",
			BankIDCode:        "GBDSC",
			AccountName:       "W Owens",
			AccountNumber:     "31926819",
			AccountNumberCode: "BBAN",
			AccountType:       0,
		},
		Debtor: Party{
			Name:              "Emelia Jane Brown",
			Address:           "10 Debtor Crescent Sourcetown NE1",
			BankID:            "203301",
			BankIDCode:        "GBDSC",
			AccountName:       "EJ Brown Black",
			AccountNumber:     "GB29XABC10161234567801",
			AccountNumberCode: "IBAN",
		},
		ChargesInformation: ChargesInformation{
			BearerCode: "SHAR",
			SenderCharges: []Charge{
				{Amount: 5.00, Currency: "GBP"},
				{Amount: 10.00, Currency: "USD"},
			},
			ReceiverCharge: Charge{
				Amount:   1.00,
				Currency: "USD",
			},
		},
		ForeignExchange: ForeignExchange{
			ContractReference: "FX123",
			ExchangeRate:      2.00000,
			OriginalAmount: Charge{
				Amount:   200.42,
				Currency: "USD",
			},
		},
		EndToEndReference:    "Wil piano Jan",
		NumericReference:     "1002001",
		PaymentID:            "123456789012345678",
		PaymentPurpouse:      "Paying for goods/services",
		PaymentSchema:        "FPS",
		PaymentType:          "Credit",
		ProcessingTime:       time.Unix(1543110881, 0),
		Reference:            "Payment for Em's piano lessons",
		SchemaPaymentType:    "ImmediatePayment",
		SchemaPaymentSubType: "InternetBanking",
		Sponsor: Party{
			AccountNumber: "56781234",
			BankID:        "123123",
			BankIDCode:    "GBDSC",
		},
	}

	assert.Equal(t, expected, payment)
}

func getJSONPayment() []byte {
	return []byte(`
	{
		"data": {
			"type": "payments",
			"id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			"attributes": {
				"amount": 100.21,
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
	}`)
}
