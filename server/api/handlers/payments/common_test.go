package payments_test

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/hmoragrega/f3-payments/server/api/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

var testClient *baloo.Client

var jsonApiContentTypePattern = `application/vnd\.api\+json`

func client() *baloo.Client {
	if testClient == nil {
		testClient = baloo.New(config.NewConfig().GetAPIEndpoint())
	}

	return testClient
}

func reloadFixtures() {
	cmd := exec.Command("make", "mongo-load-fixtures")
	cmd.Dir = "../../../../"
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Load fixtures has failed with %s\n", err)
	}
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
