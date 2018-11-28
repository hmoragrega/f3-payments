// +build unit

package payments_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/hmoragrega/f3-payments/cmd/api/handlers/payments"
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetErrorValidationFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.GetPaymentHandler(m)
	c := echoContext(http.MethodGet, "/payments/foo", strings.NewReader(""))

	m.On("Get", mock.Anything).Return(&payment.Payment{}, payment.ErrValidationFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusUnprocessableEntity, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=422, message=The payment is not valid", err.Error())
}

func TestGetServiceUnavailableError(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.GetPaymentHandler(m)
	c := echoContext(http.MethodGet, "/payments/foo", strings.NewReader(""))

	m.On("Get", mock.Anything).Return(&payment.Payment{}, payment.ErrPaymentLookup)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=There has been an error getting payment", err.Error())
}
