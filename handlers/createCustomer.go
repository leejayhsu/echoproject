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
	models.DB.Model(&models.Customer{}).Create(&customer)
	return c.JSONPretty(http.StatusCreated, customer, "  ")
}
