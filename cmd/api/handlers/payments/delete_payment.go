package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/labstack/echo"
)

// DeletePaymentHandler handle requests to delete a payment
func DeletePaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := s.Delete(c.Param("id"))

		if err != nil {
			code := http.StatusInternalServerError
			switch err {
			case payment.ErrDeleteFailed:
				code = http.StatusServiceUnavailable
			}

			return echo.NewHTTPError(code, err)
		}

		return handlers.JSONApiNoContentPretty(c)
	}
}
