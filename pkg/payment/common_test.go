package payment

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.PanicLevel)
	m.Run()
}
