package payments

// Charge represents an amount given in a specific currency
type Charge struct {
	Amount   float32 `jsonapi:"attr,amount"`
	Currency string  `jsonapi:"attr,currency"`
}
