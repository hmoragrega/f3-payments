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

func TestPatchrrorUnmarshalingPayload(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader("invalid payload"))

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=400, message=invalid character 'i' looking for beginning of value", err.Error())
}

func TestPatchErrorNotFound(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, payment.ErrPaymentNotFound)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusNotFound, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=404, message=The payment has not been found", err.Error())
}

func TestPatchErrorValidationFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, payment.ErrValidationFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=400, message=The payment is not valid", err.Error())
}

func TestPatchErrorPersistFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, payment.ErrPersistFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=The payment could not be persisted", err.Error())
}

func TestPatchErrorLookupFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, payment.ErrPaymentLookup)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=There has been an error getting payment", err.Error())
}

func TestPatchErrorMergeFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, payment.ErrMergeFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=The resulting payment is not valid", err.Error())
}

func TestPatchErrorServerError(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.PatchPaymentHandler(m)
	c := echoContext(http.MethodPatch, "/payments/foo", strings.NewReader(`{"data":{"type": "payments"}}`))

	m.On("Merge", mock.Anything, mock.Anything).Return(&payment.Payment{}, errors.New("unexpected"))

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=500, message=unexpected", err.Error())
}
