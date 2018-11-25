package payment

// Charge represents an amount given in a specific currency
type Charge struct {
	Amount   float32 `json:"amount" valid:"required" bson:"amount"`
	Currency string  `json:"currency" valid:"required" bson:"currency"`
}
