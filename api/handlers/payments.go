package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// CreatePayment creates a new a payment
func CreatePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GetPayments returns a collection of payments
func GetPayments(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GetPayment returns a single payments from the given ID
func GetPayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// UpdatePayment updates the full payment with a new payment
func UpdatePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PatchPayment updates payment attributtes
func PatchPayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// DeletePayment deletes a payment
func DeletePayment(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
