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
	mongoSession   *persistence.MongoSession
	paymentService payment.ServiceInterface
}

// NewDIC factory method to create a container
func NewDIC(c *Config) (*DIC, error) {
	ms := persistence.NewMongoSession(&c.Mongo)
	if err := ms.Connect(); err != nil {
		return nil, err
	}

	v := &validation.GoValidator{}

	return &DIC{
		mongoSession:   ms,
		paymentService: createPaymentService(c, ms, v),
	}, nil
}

// GetPaymentService gets the payment service
func (d *DIC) GetPaymentService() payment.ServiceInterface {
	return d.paymentService
}

// Clean cleans the services and dependencies
func (d *DIC) Clean() {
	d.mongoSession.Close()
}

func createPaymentService(c *Config, ms *persistence.MongoSession, v validation.Validator) payment.ServiceInterface {
	repo := persistence.NewMongoRepository(ms, createPaymentMongoEntity(), c.Database)

	return payment.NewService(repo, v)
}

func createPaymentMongoEntity() *persistence.MongoEntity {
	return persistence.NewMongoEntity(
		payment.PaymentType,
		func() interface{} { return new(payment.Payment) },
		func() interface{} { return new(payment.Collection) },
	)
}
