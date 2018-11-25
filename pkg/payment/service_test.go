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

func TestGetPaymentCorrectly(t *testing.T) {
	s, r, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(p, nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, p, result)
	assert.Nil(t, err)
}

func TestGetPaymentLookupError(t *testing.T) {
	s, r, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(nil, errors.New("foo"))

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrPaymentLookup, err)
	assert.Nil(t, result)
}

func TestGetPaymentNotFoundError(t *testing.T) {
	s, r, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(nil, nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrPaymentNotFound, err)
	assert.Nil(t, result)
}

func TestGetPaymentNotValid(t *testing.T) {
	s, r, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return("not a payment", nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrValidationFailed, err)
	assert.Nil(t, result)
}

func TestUpdateCorrectly(t *testing.T) {
	s, r, v := getServices()
	old := &Payment{ID: "foo", Amount: 10}
	new := &Payment{ID: "foo", Amount: 20}

	r.On("Get", new.ID).Return(old, nil)
	v.On("Validate", new).Return(nil)
	r.On("Persist", new).Return(nil)

	err := s.Update(new)

	assert.Nil(t, err)
}

func getServices() (ServiceInterface, *persistence.MockRepository, *validation.MockValidator) {
	r := &persistence.MockRepository{}
	v := &validation.MockValidator{}
	s := NewService(r, v)

	return s, r, v
}
