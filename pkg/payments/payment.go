package payments

// Payment represents a completed payment transaction between two parties
type Payment struct {
	ID string `jsonapi:"primary,Payment"`
}
