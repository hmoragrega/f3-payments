package payments

import (
	"bytes"
	"testing"

	"github.com/google/jsonapi"
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
				"currency": "USD"
			}
		}
	}`)
}
