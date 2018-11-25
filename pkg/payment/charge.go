package payment

// Charge represents an amount given in a specific currency
type Charge struct {
	Amount   float32 `jsonapi:"attr,amount" valid:"required" bson:"amount"`
	Currency string  `jsonapi:"attr,currency" valid:"required" bson:"currency"`
}
