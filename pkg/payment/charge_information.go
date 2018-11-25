package payment

// ChargesInformation represents the charges derived from a payment
type ChargesInformation struct {
	BearerCode     string   `jsonapi:"attr,bearer_code"`
	SenderCharges  []Charge `jsonapi:"attr,sender_charges"`
	ReceiverCharge Charge   `jsonapi:"attr,receiver_charge"`
}
