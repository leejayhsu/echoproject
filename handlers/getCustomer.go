package handlers

import (
	"echoproject/models"
	"echoproject/types"
	"net/http"

	"github.com/labstack/echo"
)

// GetCustomer creates a customer
func GetCustomer(c echo.Context) error {
	customerID := c.Param("customerID")
	var customer types.Customer
	models.DB.Where("customer_id = ?", customerID).First(&customer)
	return c.JSONPretty(http.StatusOK, customer, "  ")
}
