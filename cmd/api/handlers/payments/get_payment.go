package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/labstack/echo"
)

// GetPaymentHandler handle requests to get a payment by the id
func GetPaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {
		p, err := s.Get(c.Param("id"))

		if err != nil {
			return getGetErrorResponse(err)
		}

		return handlers.JSONApiPretty(c, http.StatusOK, p)
	}
}

func getGetErrorResponse(err error) error {
	code := http.StatusInternalServerError
	switch err {
	case payment.ErrPaymentNotFound:
		code = http.StatusNotFound
	case payment.ErrValidationFailed:
		code = http.StatusUnprocessableEntity
	case payment.ErrPaymentLookup:
		code = http.StatusServiceUnavailable
	}

	return echo.NewHTTPError(code, err)
}
