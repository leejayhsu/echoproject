package handlers

import (
	"echoproject/models"
	"net/http"

	"github.com/labstack/echo"
)

// DeleteCustomer creates a customer
func DeleteCustomer(c echo.Context) error {
	customerID := c.Param("customerID")
	models.DB.Model(&models.Customer{}).Where("customer_id = ?", customerID).Delete(&models.Customer{})
	// models.DB.Delete(&customerID) this doesn't work
	// models.DB.Where("customer_id = ?", customerID).Delete(&models.Customer{})
	// models.DB.Select("Addresses").Delete(models.Customer{}, "customer_id = ?", customerID)
	// models.DB.Where("customer_id = ?", customerID).Select("Address").Delete(&models.Customer{})
	// models.DB.Select("Addresses").Where("customer_id = ?", customerID).Delete(&models.Customer{})
	// models.DB.Select(clause.Assocations).Where("customer_id = ?", customerID).Delete(&models.Customer{})
	// models.DB.Where("customer_id = ?", customerID).Select("addresses").Delete(&models.Customer{})
	response := map[string](string){"status": customerID + " deleted"}
	return c.JSONPretty(http.StatusOK, response, "  ")
}
