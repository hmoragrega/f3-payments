package validation

// Validator interface implemented for structures that can be validated
type Validator interface {
	Validate(i interface{}) error
}
