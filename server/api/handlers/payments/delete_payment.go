package payments

import (
	"net/http"

	"github.com/labstack/echo"
)

// DeletePayment deletes a payment
func DeletePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
