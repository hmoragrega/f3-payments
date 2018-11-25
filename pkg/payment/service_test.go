package payment

import (
	"errors"
	"testing"

	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"

	"github.com/stretchr/testify/assert"
)

func TestCreateCorrectly(t *testing.T) {
	s, r, v := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(nil)
	r.On("Persist", p).Return(nil)

	err := s.Create(p)

	assert.Nil(t, err)
}

func TestCreateWithValidationError(t *testing.T) {
	s, _, v := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(errors.New("foo"))

	err := s.Create(p)

	assert.Equal(t, ErrValidationFailed, err)
}

func TestCreateWithPersitenceError(t *testing.T) {
	s, r, v := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(nil)
	r.On("Persist", p).Return(errors.New("foo"))

	err := s.Create(p)

	assert.Equal(t, ErrPersistFailed, err)
}

func getServices() (ServiceInterface, *persistence.MockRepository, *validation.MockValidator) {
	r := &persistence.MockRepository{}
	v := &validation.MockValidator{}
	s := NewService(r, v)

	return s, r, v
}
