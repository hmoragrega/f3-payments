package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hmoragrega/f3-payments/cmd/api/config"
	log "github.com/sirupsen/logrus"
)

// Wait 10 seconds before forcing the server shutdown
const timeout = 10

func main() {
	c := config.NewConfig()

	d, err := config.NewDIC(c)
	if err != nil {
		log.Fatal(err)
	}

	api := NewF3API(c, d)
	go api.Start()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := api.Shutdown(ctx); err != nil {
		api.Logger.Fatal(err)
	}
}
