package payments

import (
	"net/http"

	"github.com/labstack/echo"
)

// CreatePayment creates a new Payment
func CreatePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
