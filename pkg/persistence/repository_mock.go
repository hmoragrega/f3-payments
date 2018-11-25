package persistence

import "github.com/stretchr/testify/mock"

// MockRepository is a mocked object that implements the repository interface
type MockRepository struct {
	mock.Mock
}

// Persist persists an entity
func (r *MockRepository) Persist(i interface{}) error {
	return r.Called(i).Error(0)
}

// List returns a list of payments
func (r *MockRepository) List() (interface{}, error) {
	return nil, nil
}

// Get returns a single payment
func (r *MockRepository) Get(ID string) (interface{}, error) {
	return nil, nil
}

// Delete deletes a payment
func (r *MockRepository) Delete(ID string) error {
	return nil
}
