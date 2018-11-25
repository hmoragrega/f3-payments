package payments

// Payment represents a completed payment transaction between two parties
type Payment struct {
	ID                 string             `jsonapi:"primary,payments"`
	Amount             float32            `jsonapi:"attr,amount"`
	Currency           string             `jsonapi:"attr,currency"`
	BeneficiaryParty   PaymentParty       `jsonapi:"attr,beneficiary_party"`
	DebtorParty        PaymentParty       `jsonapi:"attr,debtor_party"`
	ChargesInformation ChargesInformation `jsonapi:"attr,charges_information"`
}
