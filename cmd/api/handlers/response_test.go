package handlers

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
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
