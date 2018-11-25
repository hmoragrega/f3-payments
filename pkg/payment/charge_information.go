package payment

// ChargesInformation represents the charges derived from a payment
type ChargesInformation struct {
	BearerCode     string   `json:"bearer_code" valid:"required,stringlength(2|100)" bson:"bearer_code"`
	SenderCharges  []Charge `json:"sender_charges" valid:"required" bson:"sender_charges"`
	ReceiverCharge Charge   `json:"receiver_charge" valid:"required" bson:"receiver_charge"`
}
