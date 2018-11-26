package config

import (
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"
)

var (
	service payment.ServiceInterface
)

// DIC a simple dependency injection container
type DIC struct {
	config         *Config
	paymentService payment.ServiceInterface
}

// NewDIC factory method to create a container
func NewDIC(c *Config) (*DIC, error) {
	ps, err := createPaymentService(c)
	if err != nil {
		return nil, err
	}

	return &DIC{
		config:         c,
		paymentService: ps,
	}, nil
}

// GetPaymentService gets the payment service
func (d *DIC) GetPaymentService() payment.ServiceInterface {
	return d.paymentService
}

func createPaymentService(c *Config) (payment.ServiceInterface, error) {
	r, err := mongoRepository(c.Mongo, createPaymentMongoEntity())
	if err != nil {
		return nil, err
	}

	return payment.NewService(r, validator()), nil
}

func createPaymentMongoEntity() *persistence.MongoEntity {
	return persistence.NewMongoEntity(
		payment.PaymentType,
		func() interface{} { return new(payment.Payment) },
		func() interface{} { return new(payment.Collection) },
	)
}

func mongoRepository(c persistence.MongoConfig, entity *persistence.MongoEntity) (persistence.Repository, error) {
	return persistence.NewMongoRepository(c, entity)
}

func validator() validation.Validator {
	return &validation.GoValidator{}
}
