package payment

// ForeignExchange represents an foreign currency exchange in a given time
type ForeignExchange struct {
	ContractReference string  `jsonapi:"attr,contract_reference"`
	ExchangeRate      float32 `jsonapi:"attr,exchange_rate"`
	OriginalAmount    Charge  `jsonapi:"attr,original_amount"`
}
