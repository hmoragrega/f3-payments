package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hmoragrega/f3-payments/pkg/logging"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const jsonIndent = "  "

var (
	// ErrInvalidJSONAPI error triggered when marshaling a struct to json:api has failed
	ErrInvalidJSONAPI = errors.New("The response could not be marshalled as json api")

	// ErrInvalidJSON error triggered when marshaling a struct to json has failed
	ErrInvalidJSON = errors.New("The response could not be marshalled as json")
)

type jsonAPIError struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type jsonAPIErrorPayload struct {
	Errors []jsonAPIError `json:"errors"`
}

// JSONApiPretty creates a valid json:api response
func JSONApiPretty(c echo.Context, code int, i interface{}) error {
	payload, err := jsonapi.Marshal(i)
	if err != nil {
		return logging.Errors(ErrInvalidJSONAPI, err)
	}

	b, err := json.MarshalIndent(payload, "", jsonIndent)
	if err != nil {
		return logging.Errors(ErrInvalidJSON, err)
	}

	return c.Blob(code, jsonapi.MediaType, b)
}

// JSONApiNoContentPretty sends an empty response
func JSONApiNoContentPretty(c echo.Context) error {
	return c.Blob(http.StatusNoContent, jsonapi.MediaType, nil)
}

// JSONApiErrorPrettyHanler http error handler
func JSONApiErrorPrettyHanler(err error, c echo.Context) {

	code := http.StatusInternalServerError

	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
	}

	jsonAPIError := &jsonAPIErrorPayload{
		Errors: []jsonAPIError{{Status: code, Detail: err.Error()}},
	}

	c.JSONPretty(code, jsonAPIError, jsonIndent)
}
