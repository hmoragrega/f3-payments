package payment

import (
	"errors"

	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"
)

var (
	// ErrPersistFailed is triggered when there is a failure persisting a payment
	ErrPersistFailed = errors.New("The payment could not be persisted")

	// ErrValidationFailed is triggered when a payment is not valid
	ErrValidationFailed = errors.New("The payment is not valid")

	// ErrPaymentNotFound is triggered when a payment is not found
	ErrPaymentNotFound = errors.New("The payment has not been found")

	// ErrPaymentLookup is triggered when a payment lookup fails
	ErrPaymentLookup = errors.New("There has been an error getting payment")

	// ErrDeleteFailed is triggered when there is a failure deleting a payment
	ErrDeleteFailed = errors.New("The payment could not be deleted")
)

// ServiceInterface payment service public API
type ServiceInterface interface {
	Create(p *Payment) error
	Update(p *Payment) error
	List() (PaymentCollection, error)
	Get(ID string) (*Payment, error)
	Delete(ID string) error
}

// Service the general payment service
type service struct {
	repo      persistence.Repository
	validator validation.Validator
}

// NewService factory method to create a payment service
func NewService(repo persistence.Repository, validator validation.Validator) ServiceInterface {
	return &service{repo, validator}
}

// Create creates a new payment
func (s *service) Create(p *Payment) error {
	return s.persist(p)
}

// Update updates an existing payment
func (s *service) Update(p *Payment) error {
	_, err := s.Get(p.ID)
	if err != nil {
		return err
	}

	return s.persist(p)
}

// List gets the collection of payments
func (s *service) List() (PaymentCollection, error) {
	i, err := s.repo.List()
	if err != nil {
		return nil, ErrPaymentLookup
	}

	l, ok := i.(PaymentCollection)
	if !ok {
		return nil, ErrValidationFailed
	}

	return l, nil
}

// Get retrieves a single payment
func (s *service) Get(ID string) (*Payment, error) {
	i, err := s.repo.Get(ID)
	if err != nil {
		return nil, ErrPaymentLookup
	}

	if i == nil {
		return nil, ErrPaymentNotFound
	}

	p, ok := i.(*Payment)
	if !ok {
		return nil, ErrValidationFailed
	}

	return p, nil
}

// Delete deletes a
func (s *service) Delete(ID string) error {
	if err := s.repo.Delete(ID); err != nil {
		return ErrDeleteFailed
	}

	return nil
}

func (s *service) persist(p *Payment) error {
	if err := s.validator.Validate(p); err != nil {
		return ErrValidationFailed
	}

	if err := s.repo.Persist(p); err != nil {
		return ErrPersistFailed
	}

	return nil
}
