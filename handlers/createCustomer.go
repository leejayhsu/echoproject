package handlers

import (
	"echoproject/models"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// Customer is a customer
type Customer struct {
	CustomerID string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	DOB        string `json:"dob"`
}

// CreateCustomer creates a customer
func CreateCustomer(c echo.Context) error {
	customerID := uuid.New().String()
	customer := new(Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	customer.CustomerID = customerID
	log.Printf("Creating customer %s", customerID)
	models.DB.Create(&customer)
	return c.JSONPretty(http.StatusCreated, customer, "  ")
}
