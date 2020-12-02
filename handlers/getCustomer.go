package handlers

import (
	berrors "echoproject/errors"
	"echoproject/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// GetCustomer creates a customer
func GetCustomer(c echo.Context) error {
	customerID := c.Param("customerID")
	var customer models.Customer
	result := models.DB.Where("customer_id = ?", customerID).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			extra := map[string](string){"status": "this is an extra"}
			err := &berrors.BError{
				Message: fmt.Sprintf("could not find customer_id %s", customerID),
				Code:    404,
				Status:  "custom error",
				Kind:    "request error",
				Extra:   extra,
			}
			return err
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "error occured while retrieving customer")
	}
	return c.JSONPretty(http.StatusOK, customer, "  ")
}
