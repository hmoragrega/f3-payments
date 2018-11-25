package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/server/api/handlers"

	"github.com/labstack/echo"
)

// ListPaymentHandler handle requests to get a list of payments
func ListPaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {
		l, err := s.List()

		if err != nil {
			code := http.StatusInternalServerError
			switch err {
			case payment.ErrValidationFailed:
				code = http.StatusUnprocessableEntity
			case payment.ErrPaymentLookup:
				code = http.StatusServiceUnavailable
			}

			return echo.NewHTTPError(code, err)
		}

		return handlers.JSONApiPretty(c, http.StatusOK, *l)
	}
}
