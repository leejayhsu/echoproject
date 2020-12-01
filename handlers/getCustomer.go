package handlers

import (
	"echoproject/models"
	"net/http"

	"github.com/labstack/echo"
)

// GetCustomer creates a customer
func GetCustomer(c echo.Context) error {
	customerID := c.Param("customerID")
	var customer models.Customer
	models.DB.Where("customer_id = ?", customerID).First(&customer)
	return c.JSONPretty(http.StatusOK, customer, "  ")
}
