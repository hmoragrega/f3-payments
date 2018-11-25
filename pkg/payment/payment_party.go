package payment

// Party represents an party involved in a payment, either the beneficiary or the debtor
type Party struct {
	Name              string `json:"name,omitempty" bson:"name"`
	Address           string `json:"address,omitempty" bson:"address"`
	BankID            string `json:"bank_id" valid:"required" bson:"bank_id"`
	BankIDCode        string `json:"bank_id_code" valid:"required" bson:"bank_id_code"`
	AccountName       string `json:"account_name,omitempty" bson:"account_name"`
	AccountNumber     string `json:"account_number" valid:"required" bson:"account_number"`
	AccountNumberCode string `json:"account_number_code,omitempty" valid:"in(BBAN|IBAN)" bson:"account_number_code"`
	AccountType       int    `json:"account_type,omitempty" bson:"account_type"`
}
