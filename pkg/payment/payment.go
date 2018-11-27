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
	ID                   string             `jsonapi:"primary,payments" valid:"uuidv4" bson:"_id"`
	Amount               float32            `jsonapi:"attr,amount" valid:"required" bson:"amount"`
	Currency             string             `jsonapi:"attr,currency" valid:"required,in(USD|GBP|EUR)" bson:"currency"`
	EndToEndReference    string             `jsonapi:"attr,end_to_end_reference" valid:"required,stringlength(2|100)" bson:"end_to_end_reference"`
	Reference            string             `jsonapi:"attr,reference" valid:"required,stringlength(2|100)"  bson:"reference"`
	NumericReference     string             `jsonapi:"attr,numeric_reference" valid:"required,stringlength(2|100)" bson:"numeric_reference"`
	PaymentID            string             `jsonapi:"attr,payment_id" valid:"required,stringlength(2|100)" bson:"payment_id"`
	PaymentPurpouse      string             `jsonapi:"attr,payment_purpose" valid:"required,stringlength(2|100)" bson:"payment_purpose"`
	PaymentSchema        string             `jsonapi:"attr,payment_scheme" valid:"required,stringlength(2|100)" bson:"payment_scheme"`
	PaymentType          string             `jsonapi:"attr,payment_type" valid:"required,stringlength(2|100)" bson:"payment_type"`
	SchemaPaymentType    string             `jsonapi:"attr,scheme_payment_type" valid:"required,stringlength(2|100)" bson:"scheme_payment_type"`
	SchemaPaymentSubType string             `jsonapi:"attr,scheme_payment_sub_type" valid:"required,stringlength(2|100)" bson:"scheme_payment_sub_type"`
	ProcessingTime       time.Time          `jsonapi:"attr,processing_time" bson:"processing_time"`
	Beneficiary          Party              `jsonapi:"attr,beneficiary_party" json:"beneficiary_party" valid:"required" bson:"beneficiary_party"`
	Debtor               Party              `jsonapi:"attr,debtor_party" valid:"required" bson:"debtor_party"`
	Sponsor              Party              `jsonapi:"attr,sponsor_party" valid:"required" bson:"sponsor_party"`
	ChargesInformation   ChargesInformation `jsonapi:"attr,charges_information" valid:"required" bson:"charges_information"`
	ForeignExchange      ForeignExchange    `jsonapi:"attr,fx" valid:"required" bson:"fx"`
}

// ForeignExchange represents an foreign currency exchange in a given time
type ForeignExchange struct {
	ContractReference string  `jsonapi:"attr,contract_reference" json:"contract_reference" valid:"required,stringlength(2|100)" bson:"contract_reference"`
	ExchangeRate      float32 `jsonapi:"attr,exchange_rate" json:"exchange_rate" valid:"required" bson:"exchange_rate"`
	OriginalAmount    Charge  `jsonapi:"attr,original_amount" json:"original_amount" valid:"required" bson:"original_amount"`
}

// Party represents an party involved in a payment, either the beneficiary or the debtor
type Party struct {
	Name              string `jsonapi:"attr,name,omitempty" json:"name,omitempty" bson:"name"`
	Address           string `jsonapi:"attr,address,omitempty" json:"address,omitempty" bson:"address"`
	BankID            string `jsonapi:"attr,bank_id" json:"bank_id" valid:"required" bson:"bank_id"`
	BankIDCode        string `jsonapi:"attr,bank_id_code" json:"bank_id_code" valid:"required" bson:"bank_id_code"`
	AccountName       string `jsonapi:"attr,account_name,omitempty"  json:"account_name,omitempty" bson:"account_name"`
	AccountNumber     string `jsonapi:"attr,account_number" json:"account_number" valid:"required" bson:"account_number"`
	AccountNumberCode string `jsonapi:"attr,account_number_code,omitempty" json:"account_number_code,omitempty" valid:"in(BBAN|IBAN)" bson:"account_number_code"`
	AccountType       int    `jsonapi:"attr,account_type,omitempty" json:"account_type,omitempty" bson:"account_type"`
}

// ChargesInformation represents the charges derived from a payment
type ChargesInformation struct {
	BearerCode     string   `jsonapi:"attr,bearer_code" json:"bearer_code" valid:"required,stringlength(2|100)" bson:"bearer_code"`
	SenderCharges  []Charge `jsonapi:"attr,sender_charges" json:"sender_charges" valid:"required" bson:"sender_charges"`
	ReceiverCharge Charge   `jsonapi:"attr,receiver_charge" json:"receiver_charge" valid:"required" bson:"receiver_charge"`
}

// Charge represents an amount given in a specific currency
type Charge struct {
	Amount   float32 `jsonapi:"attr,amount" json:"amount" valid:"required" bson:"amount"`
	Currency string  `jsonapi:"attr,currency" json:"currency" valid:"required" bson:"currency"`
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
