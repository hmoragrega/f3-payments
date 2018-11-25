package payment

import (
	"fmt"

	"github.com/google/jsonapi"
)

// Collection represents a collection of payments
type Collection []*Payment

// JSONAPILinks adds links to the json api response for the payment collection
func (p Collection) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("/%s", PaymentType),
	}
}
