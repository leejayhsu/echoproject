package models

import (
	"gorm.io/gorm"
)

// DB is global variable for db connection pool
var DB *gorm.DB

// Customer is a customer model
type Customer struct {
	gorm.Model
	CustomerID string
	FirstName  string
	LastName   string
	DOB        string
}
