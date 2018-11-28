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

func TestDeleteFailError(t *testing.T) {
	m := &PaymentServiceMock{}
	h := payments.DeletePaymentHandler(m)
	c := echoContext(http.MethodDelete, "/payments/foo", strings.NewReader(""))

	m.On("Delete", mock.Anything).Return(payment.ErrDeleteFailed)

	err := h(c)

	assert.IsType(t, &echo.HTTPError{}, err)
	assert.Equal(t, http.StatusServiceUnavailable, err.(*echo.HTTPError).Code)
	assert.Equal(t, "code=503, message=The payment could not be deleted", err.Error())
}
