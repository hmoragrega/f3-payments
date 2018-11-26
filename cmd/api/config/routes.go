package config

import (
	"github.com/hmoragrega/f3-payments/cmd/api/handlers/payments"
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/labstack/echo"
)

// RegisterRoutes registers the routes availbale in the API
func RegisterRoutes(e *echo.Echo, d *DIC) error {
	registerPaymentRoutes(e, d.GetPaymentService())

	return nil
}

// Registers all routes for the "/payments" group
func registerPaymentRoutes(e *echo.Echo, s payment.ServiceInterface) {
	g := e.Group(payment.PaymentType)

	g.GET("", payments.ListPaymentHandler(s))
	g.GET("/:id", payments.GetPaymentHandler(s))

	// Non-safe methods
	g.POST("", payments.CreatePaymentHandler(s))
	g.PUT("/:id", payments.ReplacePaymentHandler(s))
	g.PATCH("/:id", payments.PatchPaymentHandler(s))
	g.DELETE("/:id", payments.DeletePaymentHandler(s))
}
