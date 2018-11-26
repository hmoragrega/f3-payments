package payments

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/labstack/echo"
)

// CreatePaymentHandler handle requests to create a payment
func CreatePaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {

		p := new(payment.Payment)

		if err := jsonapi.UnmarshalPayload(c.Request().Body, p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		err := s.Create(p)

		if err != nil {
			code := http.StatusInternalServerError
			switch err {
			case payment.ErrValidationFailed:
				code = http.StatusBadRequest
			case payment.ErrPersistFailed:
				code = http.StatusServiceUnavailable
			}

			return echo.NewHTTPError(code, err)
		}

		return handlers.JSONApiPretty(c, http.StatusCreated, p)
	}
}
