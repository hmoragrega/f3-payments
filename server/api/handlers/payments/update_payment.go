package payments

import (
	"net/http"

	"github.com/labstack/echo"
)

// ReplacePayment replaces a payment
func ReplacePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
