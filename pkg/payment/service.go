package payment

import (
	"errors"

	log "github.com/hmoragrega/f3-payments/pkg/logging"
	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"
	"github.com/imdario/mergo"
	uuid "github.com/satori/go.uuid"
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

	// ErrMergeFailed is triggered when a payment cannot be merged into another
	ErrMergeFailed = errors.New("The payment is not valid")
)

// ServiceInterface payment service public API
type ServiceInterface interface {
	Create(p *Payment) error
	Update(ID string, p *Payment) error
	Merge(ID string, p *Payment) (*Payment, error)
	List() (*Collection, error)
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
	p.ID = uuid.NewV4().String()
	return s.persist(p)
}

// Update updates an existing payment
func (s *service) Update(ID string, p *Payment) error {
	p.ID = ID
	_, err := s.Get(p.ID)
	if err != nil {
		if err != ErrPaymentNotFound {
			return err
		}

		return s.persist(p)
	}

	return s.update(p)
}

// Merge search for a payment and updates some fields from it, returning the final payment
func (s *service) Merge(ID string, p *Payment) (*Payment, error) {
	o, err := s.Get(ID)
	if err != nil {
		return nil, err
	}

	p.ID = ID
	if err := mergo.Merge(o, p, mergo.WithOverride); err != nil {
		return nil, log.Errors(ErrMergeFailed, err)
	}

	if err := s.update(o); err != nil {
		return nil, err
	}

	return o, nil
}

// List gets the collection of payments
func (s *service) List() (*Collection, error) {
	i, err := s.repo.List()
	if err != nil {
		return nil, log.Errors(ErrPaymentLookup, err)
	}

	l, ok := i.(*Collection)
	if !ok {
		return nil, log.Errors(ErrValidationFailed)
	}

	return l, nil
}

// Get retrieves a single payment
func (s *service) Get(ID string) (*Payment, error) {
	i, err := s.repo.Get(ID)
	if err != nil {
		return nil, log.Errors(ErrPaymentLookup, err)
	}

	if i == nil {
		return nil, log.Errors(ErrPaymentNotFound)
	}

	p, ok := i.(*Payment)
	if !ok {
		return nil, log.Errors(ErrValidationFailed)
	}

	return p, nil
}

// Delete deletes a
func (s *service) Delete(ID string) error {
	if err := s.repo.Delete(ID); err != nil {
		return log.Errors(ErrDeleteFailed)
	}

	return nil
}

func (s *service) persist(p *Payment) error {
	if err := s.validator.Validate(p); err != nil {
		return log.Errors(ErrValidationFailed, err)
	}

	if err := s.repo.Persist(p); err != nil {
		return log.Errors(ErrPersistFailed, err)
	}

	return nil
}

func (s *service) update(p *Payment) error {
	if err := s.validator.Validate(p); err != nil {
		return log.Errors(ErrValidationFailed, err)
	}

	if err := s.repo.Update(p.ID, p); err != nil {
		return log.Errors(ErrPersistFailed, err)
	}

	return nil
}
