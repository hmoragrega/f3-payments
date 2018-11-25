package validation

import (
	"github.com/stretchr/testify/mock"
)

// MockValidator is a mocked object that implements the validator interface
type MockValidator struct {
	mock.Mock
}

// Validate mocks the validation of an object
func (v *MockValidator) Validate(i interface{}) error {
	return v.Called(i).Error(0)
}
