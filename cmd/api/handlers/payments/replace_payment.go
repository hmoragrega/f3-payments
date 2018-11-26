package payments

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/hmoragrega/f3-payments/pkg/payment"

	"github.com/labstack/echo"
)

// ReplacePaymentHandler handle requests to update a payment
func ReplacePaymentHandler(s payment.ServiceInterface) func(c echo.Context) error {
	return func(c echo.Context) error {

		p := new(payment.Payment)
		if err := jsonapi.UnmarshalPayload(c.Request().Body, p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		err := s.Update(c.Param("id"), p)

		if err != nil {
			return getReplaceErrorResponse(err)
		}

		return handlers.JSONApiPretty(c, http.StatusOK, p)
	}
}

func getReplaceErrorResponse(err error) error {
	code := http.StatusInternalServerError
	switch err {
	case payment.ErrValidationFailed:
		code = http.StatusBadRequest
	case payment.ErrPaymentLookup:
	case payment.ErrPersistFailed:
		code = http.StatusServiceUnavailable
	}

	return echo.NewHTTPError(code, err)
}
