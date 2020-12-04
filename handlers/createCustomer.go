package handlers

import (
	"echoproject/models"
	"net/http"

	"github.com/labstack/echo"
)

// CreateCustomer creates a customer
func CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	// must provide arg to Model to use hooks
	result := models.DB.Model(&models.Customer{}).Create(&customer)
	if result.Error != nil {
		return result.Error
	}
	return c.JSONPretty(http.StatusCreated, customer, "  ")
}
