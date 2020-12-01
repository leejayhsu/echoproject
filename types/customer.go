package types

import "time"

// Customer is a customer
type Customer struct {
	CustomerID string    `json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	DOB        string    `json:"dob"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
