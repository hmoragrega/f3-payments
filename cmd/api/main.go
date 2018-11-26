package main

import (
	"github.com/hmoragrega/f3-payments/cmd/api/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := config.NewConfig()

	d, err := config.NewDIC(c)
	if err != nil {
		log.Fatal(err)
	}

	api := NewF3API(c, d)
	api.Start()
}
