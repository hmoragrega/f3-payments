package payments_test

import (
	"fmt"

	"github.com/hmoragrega/f3-payments/server/api/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

var testClient *baloo.Client

func client() *baloo.Client {
	if testClient == nil {
		testClient = baloo.New(config.NewConfig().GetAPIEndpoint())
	}

	return testClient
}

func getErrorResponse(code int, detail string) string {
	return fmt.Sprintf(`{
		"errors": [
			{
				"status": %d,
				"detail": "%s"
			}
		]
	}`, code, detail)
}
