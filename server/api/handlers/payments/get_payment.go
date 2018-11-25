package payments

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetPayment returns a single payments from the given ID
func GetPayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
