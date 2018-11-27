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

// Update updates an entity
func (r *MockRepository) Update(ID string, i interface{}) error {
	return r.Called(ID, i).Error(0)
}

// List returns a collection of entities
func (r *MockRepository) List() (interface{}, error) {
	args := r.Called()
	return args.Get(0), args.Error(1)
}

// Get retrieves a single entity by the ID
func (r *MockRepository) Get(ID string) (interface{}, error) {
	args := r.Called(ID)
	return args.Get(0), args.Error(1)
}

// Delete deletes an entity by the ID
func (r *MockRepository) Delete(ID string) error {
	return r.Called(ID).Error(0)
}

// DeleteAll deletes all entitis in the repository
func (r *MockRepository) DeleteAll() error {
	return r.Called().Error(0)
}
