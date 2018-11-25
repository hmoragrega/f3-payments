package payment

import (
	"fmt"

	"github.com/google/jsonapi"
)

// PaymentCollection represents a collection of payments
type PaymentCollection []*Payment

// JSONAPILinks adds links to the json api response for the payment collection
func (p PaymentCollection) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("/%s", PaymentType),
	}
}
