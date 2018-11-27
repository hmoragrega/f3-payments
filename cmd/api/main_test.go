// +build api

package main

import (
	"flag"
	"fmt"
	"testing"
)

var systemTest *bool

func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
}

func TestAPI(t *testing.T) {
	if *systemTest {
		fmt.Println("Running in test mode")
		main()
	}
}
