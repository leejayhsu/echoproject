package handlers

import (
	"echoproject/models"
	"net/http"

	"github.com/labstack/echo"
)

// DeleteCustomer creates a customer
func DeleteCustomer(c echo.Context) error {
	customerID := c.Param("customerID")
	models.DB.Where("customer_id = ?", customerID).Delete(&models.Customer{})
	response := map[string](string){"status": customerID + " deleted"}
	return c.JSONPretty(http.StatusOK, response, "  ")
}
