package handlers

import (
	"encoding/json"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidJSONAPI error triggered when marshaling a struct to json:api has failed
	ErrInvalidJSONAPI = errors.New("The response could not be marshalled as json api")

	// ErrInvalidJSON error triggered when marshaling a struct to json has failed
	ErrInvalidJSON = errors.New("The response could not be marshalled as json")
)

// JSONApiPretty creates a valid json:api response
func JSONApiPretty(c echo.Context, code int, i interface{}) error {
	payload, err := jsonapi.Marshal(i)
	if err != nil {
		log.Error(ErrInvalidJSONAPI, err)
		return ErrInvalidJSONAPI
	}

	b, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Error(ErrInvalidJSON, err)
		return ErrInvalidJSON
	}

	return c.Blob(code, jsonapi.MediaType, b)
}
