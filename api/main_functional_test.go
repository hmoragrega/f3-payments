// +build !unit functional

package main

import (
	"fmt"
	"net/http"
	"testing"

	baloo "gopkg.in/h2non/baloo.v3"
)

var client = getClient()

func TestGetPaymentCollection(t *testing.T) {
	client.Get("/payments").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestGetPayment(t *testing.T) {
	client.Get("/payments/foo").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestCreatePayment(t *testing.T) {
	client.Post("/payments").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestUpdatePayment(t *testing.T) {
	client.Put("/payments/foo").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestPatchPayment(t *testing.T) {
	client.Patch("/payments/foo").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestDeletePayment(t *testing.T) {
	client.Delete("/payments/foo").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func getClient() *baloo.Client {
	initConfig()
	return baloo.New(getEndpoint())
}

func getEndpoint() string {
	return fmt.Sprintf("http://%s", getAddress())
}
