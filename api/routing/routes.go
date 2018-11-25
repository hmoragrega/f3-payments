package routing

import (
	"net/http"

	"github.com/labstack/echo"
)

// RegisterRoutes registers the routes availbale in the API
func RegisterRoutes(e *echo.Echo) {
	registerPaymentRoutes(e)
}

func registerPaymentRoutes(e *echo.Echo) {
	//g := e.Group("payments")

	e.GET("payments", dummyHandler)
	e.GET("/payments/:id", dummyHandler)

	// Non-safe methods
	e.POST("/payments", dummyHandler)
	e.PUT("/payments", dummyHandler)
	e.PATCH("/payments", dummyHandler)
	e.PUT("/payments", dummyHandler)
	e.DELETE("/payments/:id", dummyHandler)
}

func dummyHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
