package main

import (
	"github.com/hmoragrega/f3-payments/server/api"
	"github.com/hmoragrega/f3-payments/server/api/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := config.NewConfig()

	d, err := config.NewDIC(c)
	if err != nil {
		log.Fatal(err)
	}

	a := api.NewF3API(c, d)
	a.Start()
}
