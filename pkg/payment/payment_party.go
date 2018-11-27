package payment

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
