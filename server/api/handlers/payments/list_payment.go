package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/server/api/handlers"
	"github.com/labstack/echo"
)

// GetPayments returns a collection of payments
func GetPayments(c echo.Context) error {
	return handlers.JSONApiPretty(c, http.StatusOK, payment.PaymentCollection{{ID: "foo"}})
}
