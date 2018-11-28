package handlers

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.PanicLevel)
	m.Run()
}

func TestJSONApiPrettyJSONApiMarshalError(t *testing.T) {
	e := echo.New()
	c := e.NewContext(nil, nil)

	err := JSONApiPretty(c, http.StatusOK, `invalid json:api`)

	assert.Equal(t, ErrInvalidJSONAPI, err)
}

func TestGetCodeAndMessageUnkownError(t *testing.T) {
	err := errors.New("foo")

	code, msg := getCodeAndMessage(err)

	assert.Equal(t, http.StatusInternalServerError, code)
	assert.Equal(t, "foo", msg)
}

func TestGetCodeAndMessageEchoHTTPError(t *testing.T) {
	err := echo.NewHTTPError(http.StatusBadRequest, "foo")

	code, msg := getCodeAndMessage(err)

	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, "foo", msg)
}
