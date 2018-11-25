package payment

// Party represents an party involved in a payment, either the beneficiary or the debtor
type Party struct {
	Name              string `jsonapi:"attr,name,omitempty"`
	Address           string `jsonapi:"attr,address,omitempty"`
	BankID            string `jsonapi:"attr,bank_id" valid:"required"`
	BankIDCode        string `jsonapi:"attr,bank_id_code" valid:"required"`
	AccountName       string `jsonapi:"attr,account_name,omitempty"`
	AccountNumber     string `jsonapi:"attr,account_number" valid:"required"`
	AccountNumberCode string `jsonapi:"attr,account_number_code,omitempty" valid:"in(BBAN,IBAN)"`
	AccountType       int    `jsonapi:"attr,account_type,omitempty"`
}
