package handlers

import (
	"echoproject/models"
	"echoproject/types"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// CreateCustomer creates a customer
func CreateCustomer(c echo.Context) error {
	customerID := uuid.New().String()
	customer := new(types.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	customer.CustomerID = customerID
	log.Printf("Creating customer %s", customerID)
	models.DB.Create(&customer)
	return c.JSONPretty(http.StatusCreated, customer, "  ")
}
