package payment

import (
	"fmt"
	"time"

	"github.com/google/jsonapi"
)

const (
	// PaymentType the type identifier for the payment entity
	PaymentType = "payments"

	// PaymentVersion the payment structure version
	PaymentVersion = "1.0"

	// PaymentOrganizationID the payment organization id
	PaymentOrganizationID = "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb"
)

// Payment represents a completed payment transaction between two parties
type Payment struct {
	ID                   string             `jsonapi:"primary,payments" valid:"required,uuidv4"`
	Amount               float32            `jsonapi:"attr,amount" valid:"required"`
	Currency             string             `jsonapi:"attr,currency" valid:"required,in(USD,GBP,EUR)"`
	EndToEndReference    string             `jsonapi:"attr,end_to_end_reference" valid:"required,length(3)"`
	Reference            string             `jsonapi:"attr,reference" valid:"required,length(3)"`
	NumericReference     string             `jsonapi:"attr,numeric_reference" valid:"required,length(3)"`
	PaymentID            string             `jsonapi:"attr,payment_id" valid:"required,length(3)"`
	PaymentPurpouse      string             `jsonapi:"attr,payment_purpose" valid:"required,length(3)"`
	PaymentSchema        string             `jsonapi:"attr,payment_scheme" valid:"required,length(3)"`
	PaymentType          string             `jsonapi:"attr,payment_type" valid:"required,length(3)"`
	SchemaPaymentType    string             `jsonapi:"attr,scheme_payment_type" valid:"required,length(3)"`
	SchemaPaymentSubType string             `jsonapi:"attr,scheme_payment_sub_type" valid:"required,length(3)"`
	ProcessingTime       time.Time          `jsonapi:"attr,processing_time" valid:"required"`
	Beneficiary          Party              `jsonapi:"attr,beneficiary_party" valid:"required"`
	Debtor               Party              `jsonapi:"attr,debtor_party" valid:"required"`
	Sponsor              Party              `jsonapi:"attr,sponsor_party" valid:"required"`
	ChargesInformation   ChargesInformation `jsonapi:"attr,charges_information" valid:"required"`
	ForeignExchange      ForeignExchange    `jsonapi:"attr,fx" valid:"required"`
}

// JSONAPILinks adds links to the json api response of the payment
func (p *Payment) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("/%s/%s", PaymentType, p.ID),
	}
}

// JSONAPIMeta adds metadata to each payment response
func (p *Payment) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"version":         PaymentVersion,
		"organization_id": PaymentOrganizationID,
	}
}
