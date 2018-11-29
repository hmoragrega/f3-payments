package merge

import (
	"github.com/stretchr/testify/mock"
)

// MockMerger is a mocked object that implements the merger interface
type MockMerger struct {
	mock.Mock
}

// Merge two entities
func (r *MockMerger) Merge(dst, src interface{}) error {
	return r.Called(dst, src).Error(0)
}
