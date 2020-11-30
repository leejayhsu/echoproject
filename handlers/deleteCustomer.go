package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// DeleteCustomer creates a customer
func DeleteCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
