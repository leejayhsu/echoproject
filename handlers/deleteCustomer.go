package handlers

import (
	"echoproject/models"
	"echoproject/types"
	"net/http"

	"github.com/labstack/echo"
)

// DeleteCustomer creates a customer
func DeleteCustomer(c echo.Context) error {
	customerID := c.Param("customerID")

	models.DB.Where("customer_id = ?", customerID).Delete(&models.Customer{})

	// return deleted customer data
	// can be combined into delete query using raw query
	// or just get rid of this and return a string or something
	var customer types.Customer
	models.DB.Where("customer_id = ?", customerID).First(&customer)
	return c.JSONPretty(http.StatusOK, customer, "  ")
}
