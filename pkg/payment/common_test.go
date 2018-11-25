package payment

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestMain(t *testing.M) {
	logger := logrus.New()
	logger.Out = ioutil.Discard
}
