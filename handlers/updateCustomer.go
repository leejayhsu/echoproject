package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// UpdateCustomer creates a customer
func UpdateCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
