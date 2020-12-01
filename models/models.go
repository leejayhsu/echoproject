package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DB is global variable for db connection pool
var DB *gorm.DB

//Base is same as gorm.Model except it uses uuid as pk instead of int
type Base struct {
	UUID      uuid.UUID      `gorm:"type:uuid;primaryKey;" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_a"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Customer is a customer model
type Customer struct {
	CustomerID uuid.UUID `gorm:"type:uuid" json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	DOB        string    `json:"dob"`
	Base
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New()
	return
}

// BeforeSave will set a CustomerID uuid
func (c *Customer) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("executing BeforeSave hook")
	c.CustomerID = uuid.New()
	return
}
