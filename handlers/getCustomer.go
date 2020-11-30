package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetCustomer creates a customer
func GetCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
