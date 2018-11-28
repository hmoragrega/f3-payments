package config

import (
	"github.com/hmoragrega/f3-payments/pkg/merge"
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/hmoragrega/f3-payments/pkg/validation"
)

var (
	service payment.ServiceInterface
)

// DIC a simple dependency injection container
type DIC struct {
	paymentRepository *persistence.MongoRepository
	mongoSession      *persistence.MongoSession
	paymentService    payment.ServiceInterface
}

// NewDIC factory method to create a container
func NewDIC(c *Config) (*DIC, error) {
	ms := persistence.NewMongoSession(&c.Mongo)
	if err := ms.Connect(); err != nil {
		return nil, err
	}

	pr := persistence.NewMongoRepository(ms, createPaymentMongoEntity(), c.Mongo.Database)
	v := &validation.GoValidator{}
	m := &merge.MergoMerger{}

	return &DIC{
		mongoSession:      ms,
		paymentRepository: pr,
		paymentService:    payment.NewService(pr, v, m),
	}, nil
}

// GetPaymentService gets the payment service
func (d *DIC) GetPaymentService() payment.ServiceInterface {
	return d.paymentService
}

// GetPaymentRepository gets the payment repository
func (d *DIC) GetPaymentRepository() persistence.Repository {
	return d.paymentRepository
}

// Clean cleans the services and dependencies
func (d *DIC) Clean() {
	d.mongoSession.Close()
}

func createPaymentMongoEntity() *persistence.MongoEntity {
	return persistence.NewMongoEntity(
		payment.PaymentType,
		func() interface{} { return new(payment.Payment) },
		func() interface{} { return new(payment.Collection) },
	)
}
