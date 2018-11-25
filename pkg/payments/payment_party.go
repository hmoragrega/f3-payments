package payments

// PaymentParty represents an party involved in a payment, either the beneficiary or the debtor
type PaymentParty struct {
	Name              string `jsonapi:"attr,name"`
	Address           string `jsonapi:"attr,address"`
	BankID            string `jsonapi:"attr,bank_id"`
	BankIDCode        string `jsonapi:"attr,bank_id_code"`
	AccountName       string `jsonapi:"attr,account_name"`
	AccountNumber     string `jsonapi:"attr,account_number"`
	AccountNumberCode string `jsonapi:"attr,account_number_code"`
	AccountType       int    `jsonapi:"attr,account_type,omitempty"`
}
