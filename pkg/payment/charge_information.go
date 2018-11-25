package payment

// ChargesInformation represents the charges derived from a payment
type ChargesInformation struct {
	BearerCode     string   `jsonapi:"attr,bearer_code" json:"bearer_code" valid:"required,stringlength(2|100)" bson:"bearer_code"`
	SenderCharges  []Charge `jsonapi:"attr,sender_charges" json:"sender_charges" valid:"required" bson:"sender_charges"`
	ReceiverCharge Charge   `jsonapi:"attr,receiver_charge" json:"receiver_charge" valid:"required" bson:"receiver_charge"`
}
