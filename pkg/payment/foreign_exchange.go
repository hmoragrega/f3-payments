package payment

// ForeignExchange represents an foreign currency exchange in a given time
type ForeignExchange struct {
	ContractReference string  `jsonapi:"attr,contract_reference" valid:"required,length(3)" bson:"contract_reference"`
	ExchangeRate      float32 `jsonapi:"attr,exchange_rate" valid:"required" bson:"exchange_rate"`
	OriginalAmount    Charge  `jsonapi:"attr,original_amount" valid:"required" bson:"original_amount"`
}
