package payments

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/labstack/echo"
)

// PatchPaymentHandler handle requests to pacth a payment
func PatchPaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {

		p := new(payment.Payment)

		if err := jsonapi.UnmarshalPayload(c.Request().Body, p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		o, err := s.Merge(c.Param("id"), p)

		if err != nil {
			return getPatchErrorResponse(err)
		}

		return handlers.JSONApiPretty(c, http.StatusOK, o)
	}
}

func getPatchErrorResponse(err error) error {
	code := http.StatusInternalServerError
	switch err {
	case payment.ErrPaymentNotFound:
		code = http.StatusNotFound
	case payment.ErrValidationFailed:
		code = http.StatusBadRequest
	case payment.ErrPaymentLookup:
	case payment.ErrMergeFailed:
	case payment.ErrPersistFailed:
		code = http.StatusServiceUnavailable
	}

	return echo.NewHTTPError(code, err)
}
