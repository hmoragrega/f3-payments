// +build !unit functional

package main

import (
	"fmt"
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var client = getBalooClient()

func getBalooClient() *baloo.Client {
	initConfig()
	return baloo.New(fmt.Sprintf("http://%s", getAddress()))
}

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

func TestPutPayment(t *testing.T) {
	client.Put("/payments").
		Expect(t).
		Status(http.StatusOK).
		BodyEquals("ok").
		Done()
}

func TestUpdatePayment(t *testing.T) {
	client.Patch("/payments").
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
