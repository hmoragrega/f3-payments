// +build unit

package payments_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/hmoragrega/f3-payments/cmd/api/handlers/payments"
	"github.com/hmoragrega/f3-payments/pkg/payment"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReplaceErrorUnmarshalingPayload(t *testing.T) {
	h := payments.ReplacePaymentHandler(&PaymentServiceMock{})
	c := echoContext(http.MethodPut, "/payments/foo", strings.NewReader("invalid payload"))

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=400, message=invalid character 'i' looking for beginning of value", err.Error())
}

func TestReplaceErrorValidationFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.ReplacePaymentHandler(m)
	c := echoContext(http.MethodPut, "/payments/foo", strings.NewReader(`{"data": {"type": "payments"}}`))

	m.On("Update", mock.Anything, mock.Anything).Return(payment.ErrValidationFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=400, message=The payment is not valid", err.Error())
}

func TestReplaceErrorPersistFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.ReplacePaymentHandler(m)
	c := echoContext(http.MethodPut, "/payments/foo", strings.NewReader(`{"data": {"type": "payments"}}`))

	m.On("Update", mock.Anything, mock.Anything).Return(payment.ErrPersistFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=The payment could not be persisted", err.Error())
}

func TestReplaceErrorLookupFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.ReplacePaymentHandler(m)
	c := echoContext(http.MethodPut, "/payments/foo", strings.NewReader(`{"data": {"type": "payments"}}`))

	m.On("Update", mock.Anything, mock.Anything).Return(payment.ErrPaymentLookup)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=There has been an error getting payment", err.Error())
}

func TestReplaceErrorServerError(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.ReplacePaymentHandler(m)
	c := echoContext(http.MethodPut, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Update", mock.Anything, mock.Anything).Return(errors.New("unexpected"))

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=500, message=unexpected", err.Error())
}
