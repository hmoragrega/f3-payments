package payments

import (
	"net/http"

	"github.com/labstack/echo"
)

// UpdatePayment updates payment attributtes
func UpdatePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
