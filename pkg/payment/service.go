package payment

import (
	"errors"

	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"
)

var (
	// ErrPersistFailed is triggered when there is a failure with the repository
	ErrPersistFailed = errors.New("The payment could not be persisted")

	// ErrValidationFailed is triggered when a payment is not valid
	ErrValidationFailed = errors.New("The payment is not valid")
)

// ServiceInterface payment service public API
type ServiceInterface interface {
	Create(p *Payment) error
	Update(p *Payment) error
	List(string) (PaymentCollection, error)
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
	if err := s.validator.Validate(p); err != nil {
		return ErrValidationFailed
	}

	if err := s.repo.Persist(p); err != nil {
		return ErrPersistFailed
	}

	return nil
}

// Update updates an existing payment
func (s *service) Update(p *Payment) error {
	return nil
}

// List gets the collection of payments
func (s *service) List(string) (PaymentCollection, error) {
	return nil, nil
}

// Get retrieves a single payment
func (s *service) Get(ID string) (*Payment, error) {
	return nil, nil
}

// Delete deletes a
func (s *service) Delete(ID string) error {
	return nil
}
