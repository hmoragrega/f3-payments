// +build unit

package payment

import (
	"errors"
	"testing"

	"github.com/hmoragrega/f3-payments/pkg/merge"
	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"

	"github.com/stretchr/testify/assert"
)

func TestCreateCorrectly(t *testing.T) {
	s, r, v, _ := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(nil)
	r.On("Persist", p).Return(nil)

	err := s.Create(p)

	assert.Nil(t, err)
}

func TestCreateWithValidationError(t *testing.T) {
	s, _, v, _ := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(errors.New("foo"))

	err := s.Create(p)

	assert.Equal(t, ErrValidationFailed, err)
}

func TestCreateWithPersitenceError(t *testing.T) {
	s, r, v, _ := getServices()
	p := &Payment{ID: "foo"}

	v.On("Validate", p).Return(nil)
	r.On("Persist", p).Return(errors.New("foo"))

	err := s.Create(p)

	assert.Equal(t, ErrPersistFailed, err)
}

func TestGetPaymentCorrectly(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(p, nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, p, result)
	assert.Nil(t, err)
}

func TestGetPaymentLookupError(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(nil, errors.New("foo"))

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrPaymentLookup, err)
	assert.Nil(t, result)
}

func TestGetPaymentNotFoundError(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return(nil, nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrPaymentNotFound, err)
	assert.Nil(t, result)
}

func TestGetPaymentNotValid(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Get", p.ID).Return("not a payment", nil)

	result, err := s.Get(p.ID)

	assert.Equal(t, ErrValidationFailed, err)
	assert.Nil(t, result)
}

func TestMergeCorrectly(t *testing.T) {
	s, r, v, m := getServices()
	old := &Payment{ID: "foo", Amount: 10}
	new := &Payment{ID: "foo", Amount: 20, Reference: "ref"}

	r.On("Get", new.ID).Return(old, nil)
	m.On("Merge", old, new).Return(nil)
	v.On("Validate", old).Return(nil)
	r.On("Update", old.ID, old).Return(nil)

	result, err := s.Merge(new.ID, new)

	assert.Equal(t, old, result)
	assert.Nil(t, err)
}

func TestMergeUpdateError(t *testing.T) {
	s, r, v, m := getServices()
	old := &Payment{ID: "foo", Amount: 10}
	new := &Payment{ID: "foo", Amount: 20, Reference: "ref"}

	r.On("Get", new.ID).Return(old, nil)
	m.On("Merge", old, new).Return(nil)
	v.On("Validate", old).Return(nil)
	r.On("Update", old.ID, old).Return(errors.New("foo"))

	result, err := s.Merge(new.ID, new)

	assert.Equal(t, ErrPersistFailed, err)
	assert.Nil(t, result)
}

func TestMergeValidationError(t *testing.T) {
	s, r, v, m := getServices()
	old := &Payment{ID: "foo", Amount: 10}
	new := &Payment{ID: "foo", Amount: 20, Reference: "ref"}

	r.On("Get", new.ID).Return(old, nil)
	m.On("Merge", old, new).Return(nil)
	v.On("Validate", old).Return(errors.New("foo"))

	result, err := s.Merge(new.ID, new)

	assert.Equal(t, ErrValidationFailed, err)
	assert.Nil(t, result)
}

func TestMergeNotFoundError(t *testing.T) {
	s, r, _, _ := getServices()
	new := &Payment{ID: "foo", Amount: 20, Reference: "ref"}

	r.On("Get", new.ID).Return(nil, nil)

	result, err := s.Merge(new.ID, new)

	assert.Equal(t, ErrPaymentNotFound, err)
	assert.Nil(t, result)
}

func TestMergeErrorMerging(t *testing.T) {
	s, r, _, m := getServices()
	old := &Payment{}
	new := &Payment{ID: "foo", Amount: 20, Reference: "ref"}

	r.On("Get", new.ID).Return(old, nil)
	m.On("Merge", old, new).Return(errors.New("Merge failed"))

	result, err := s.Merge(new.ID, new)

	assert.Equal(t, ErrMergeFailed, err)
	assert.Nil(t, result)
}

func TestUpdateReplacingCorrectly(t *testing.T) {
	s, r, v, _ := getServices()
	old := &Payment{ID: "foo", Amount: 10}
	new := &Payment{ID: "foo", Amount: 20}

	r.On("Get", new.ID).Return(old, nil)
	v.On("Validate", new).Return(nil)
	r.On("Update", new.ID, new).Return(nil)

	err := s.Update(new.ID, new)

	assert.Nil(t, err)
}

func TestUpdateCreatingCorrectly(t *testing.T) {
	s, r, v, _ := getServices()
	new := &Payment{ID: "foo", Amount: 20}

	r.On("Get", new.ID).Return(nil, nil)
	v.On("Validate", new).Return(nil)
	r.On("Persist", new).Return(nil)

	err := s.Update(new.ID, new)

	assert.Nil(t, err)
}

func TestUpdateWithError(t *testing.T) {
	s, r, _, _ := getServices()
	new := &Payment{ID: "foo", Amount: 20}

	r.On("Get", new.ID).Return(nil, errors.New("foo"))

	err := s.Update(new.ID, new)

	assert.Equal(t, ErrPaymentLookup, err)
}

func TestListPaymentsCorrectly(t *testing.T) {
	s, r, _, _ := getServices()
	c := &Collection{{ID: "foo"}}

	r.On("List").Return(c, nil)

	result, err := s.List()

	assert.Equal(t, c, result)
	assert.Nil(t, err)
}

func TestListPaymentsLookupError(t *testing.T) {
	s, r, _, _ := getServices()
	c := &Collection{{ID: "foo"}}

	r.On("List").Return(c, errors.New("foo"))

	result, err := s.List()

	assert.Equal(t, ErrPaymentLookup, err)
	assert.Nil(t, result)
}

func TestListPaymentsNotValid(t *testing.T) {
	s, r, _, _ := getServices()

	r.On("List").Return("not a collection", nil)

	result, err := s.List()

	assert.Equal(t, ErrValidationFailed, err)
	assert.Nil(t, result)
}

func TestDeletePaymentCorrectly(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Delete", p.ID).Return(nil)

	err := s.Delete(p.ID)

	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	s, r, _, _ := getServices()
	p := &Payment{ID: "foo"}

	r.On("Delete", p.ID).Return(errors.New("foo"))

	err := s.Delete(p.ID)

	assert.Equal(t, ErrDeleteFailed, err)
}

func getServices() (
	ServiceInterface,
	*persistence.MockRepository,
	*validation.MockValidator,
	*merge.MockMerger,
) {
	r := &persistence.MockRepository{}
	v := &validation.MockValidator{}
	m := &merge.MockMerger{}
	s := NewService(r, v, m)

	return s, r, v, m
}
