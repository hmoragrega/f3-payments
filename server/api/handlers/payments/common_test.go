package payments_test

import (
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
