package logging

import (
	log "github.com/sirupsen/logrus"
)

// Errors logs all the errors given and returns the first one
func Errors(err error, context ...error) error {
	log.Error(err, context)

	return err
}
