package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/server/api/handlers"

	"github.com/labstack/echo"
)

// GetPaymentHandler handler to get a single payment
type GetPaymentHandler struct {
	s payment.ServiceInterface
}

// NewGetPaymentHandler Factory method to create the get handler
func NewGetPaymentHandler(s payment.ServiceInterface) *GetPaymentHandler {
	return &GetPaymentHandler{s}
}

// Handle returns a single payments from the given ID
func (h *GetPaymentHandler) Handle(c echo.Context) error {

	p, err := h.s.Get(c.Param("id"))

	switch err {
	case payment.ErrPaymentNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err)
	case payment.ErrValidationFailed:
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	case payment.ErrPaymentLookup:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return handlers.JSONApiPretty(c, http.StatusOK, p)
}
