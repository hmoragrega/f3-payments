package routing

import (
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/server/api/handlers/payments"
	"github.com/labstack/echo"
)

// RegisterRoutes registers the routes availbale in the API
func RegisterRoutes(e *echo.Echo) {
	registerPaymentRoutes(e)
}

// Registers all routes for the "/payments" group
func registerPaymentRoutes(e *echo.Echo) {
	g := e.Group(payment.PaymentType)

	g.GET("", payments.GetPayments)
	g.GET("/:id", payments.GetPayment)

	// Non-safe methods
	g.POST("", payments.CreatePayment)
	g.PUT("/:id", payments.ReplacePayment)
	g.PATCH("/:id", payments.UpdatePayment)
	g.DELETE("/:id", payments.DeletePayment)
}
