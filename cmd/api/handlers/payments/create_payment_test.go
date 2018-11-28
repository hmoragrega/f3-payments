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

func TestCreateErrorUnmarshalingPayload(t *testing.T) {
	h := payments.CreatePaymentHandler(&PaymentServiceMock{})
	c := echoContext(http.MethodPost, "/payments", strings.NewReader("invalid payload"))

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=400, message=invalid character 'i' looking for beginning of value", err.Error())
}

func TestCreateErrorPersistFailed(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.CreatePaymentHandler(m)
	c := echoContext(http.MethodPost, "/payments", strings.NewReader(`{"data": {"type": "payments"}}`))

	m.On("Create", mock.Anything).Return(payment.ErrPersistFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=The payment could not be persisted", err.Error())
}
