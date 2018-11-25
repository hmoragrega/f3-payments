package payments

import (
	"bytes"
	"testing"

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
		BeneficiaryParty: PaymentParty{
			Name:              "Wilfred Jeremiah Owens",
			Address:           "1 The Beneficiary Localtown SE2",
			BankID:            "403000",
			BankIDCode:        "GBDSC",
			AccountName:       "W Owens",
			AccountNumber:     "31926819",
			AccountNumberCode: "BBAN",
			AccountType:       0,
		},
		DebtorParty: PaymentParty{
			Name:              "Emelia Jane Brown",
			Address:           "10 Debtor Crescent Sourcetown NE1",
			BankID:            "203301",
			BankIDCode:        "GBDSC",
			AccountName:       "EJ Brown Black",
			AccountNumber:     "GB29XABC10161234567801",
			AccountNumberCode: "IBAN",
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
				}
			}
		}
	}`)
}
