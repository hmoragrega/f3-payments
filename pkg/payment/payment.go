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
	ID                   string             `jsonapi:"primary,payments" valid:"required,uuidv4" bson:"_id"`
	Amount               float32            `jsonapi:"attr,amount" valid:"required" bson:"amount"`
	Currency             string             `jsonapi:"attr,currency" valid:"required,in(USD,GBP,EUR)" bson:"currency"`
	EndToEndReference    string             `jsonapi:"attr,end_to_end_reference" valid:"required,length(3)" bson:"end_to_end_reference"`
	Reference            string             `jsonapi:"attr,reference" valid:"required,length(3)"  bson:"reference"`
	NumericReference     string             `jsonapi:"attr,numeric_reference" valid:"required,length(3)" bson:"numeric_reference"`
	PaymentID            string             `jsonapi:"attr,payment_id" valid:"required,length(3)" bson:"payment_id"`
	PaymentPurpouse      string             `jsonapi:"attr,payment_purpose" valid:"required,length(3)" bson:"payment_purpose"`
	PaymentSchema        string             `jsonapi:"attr,payment_scheme" valid:"required,length(3)" bson:"payment_scheme"`
	PaymentType          string             `jsonapi:"attr,payment_type" valid:"required,length(3)" bson:"payment_type"`
	SchemaPaymentType    string             `jsonapi:"attr,scheme_payment_type" valid:"required,length(3)" bson:"scheme_payment_type"`
	SchemaPaymentSubType string             `jsonapi:"attr,scheme_payment_sub_type" valid:"required,length(3)" bson:"scheme_payment_sub_type"`
	ProcessingTime       time.Time          `jsonapi:"attr,processing_time" valid:"required" bson:"processing_time"`
	Beneficiary          Party              `jsonapi:"attr,beneficiary_party" valid:"required" bson:"beneficiary_party"`
	Debtor               Party              `jsonapi:"attr,debtor_party" valid:"required" bson:"debtor_party"`
	Sponsor              Party              `jsonapi:"attr,sponsor_party" valid:"required" bson:"sponsor_party"`
	ChargesInformation   ChargesInformation `jsonapi:"attr,charges_information" valid:"required" bson:"charges_information"`
	ForeignExchange      ForeignExchange    `jsonapi:"attr,fx" valid:"required" bson:"fx"`
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
