package routing

import (
	"github.com/hmoragrega/f3-payments/api/handlers"
	"github.com/labstack/echo"
)

// RegisterRoutes registers the routes availbale in the API
func RegisterRoutes(e *echo.Echo) {
	registerPaymentRoutes(e)
}

func registerPaymentRoutes(e *echo.Echo) {
	//g := e.Group("payments")

	e.GET("payments", handlers.GetPayments)
	e.GET("/payments/:id", handlers.GetPayment)

	// Non-safe methods
	e.POST("/payments", handlers.CreatePayment)
	e.PUT("/payments/:id", handlers.UpdatePayment)
	e.PATCH("/payments/:id", handlers.PatchPayment)
	e.DELETE("/payments/:id", handlers.DeletePayment)
}
