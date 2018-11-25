package payments

import (
	"net/http"

	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/hmoragrega/f3-payments/server/api/handlers"

	"github.com/labstack/echo"
)

// ListPaymentHandler handler to get a collection of payments
type ListPaymentHandler struct {
	s payment.ServiceInterface
}

// NewListPaymentHandler Factory method to create the list handler
func NewListPaymentHandler(s payment.ServiceInterface) *ListPaymentHandler {
	return &ListPaymentHandler{s}
}

// Handle returns all the paymentscollection of payments
func (h *ListPaymentHandler) Handle(c echo.Context) error {
	l, err := h.s.List()

	switch err {
	case payment.ErrValidationFailed:
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	case payment.ErrPaymentLookup:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return handlers.JSONApiPretty(c, http.StatusOK, *l)
}
