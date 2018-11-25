package main

import (
	"github.com/hmoragrega/f3-payments/server/api"
)

func main() {
	a := api.NewF3API()
	a.Start()
}
