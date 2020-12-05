package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DB is global variable for db connection pool
var DB *gorm.DB

//Base is same as gorm.Model except it uses uuid as pk instead of int
type Base struct {
	UUID      uuid.UUID      `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Customer is a customer model
type Customer struct {
	CustomerID uuid.UUID `gorm:"type:uuid;unique:true" json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	DOB        string    `json:"dob"`
	Addresses  []Address `gorm:"foreignKey:CustomerID;references:CustomerID;" json:"addresses"`
	Base
}

// Address is an address model
type Address struct {
	AddressID  uuid.UUID `gorm:"type:uuid" json:"address_id"`
	Street1    string    `json:"street1"`
	Street2    string    `json:"street2"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Zip        string    `json:"zip"`
	Country    string    `json:"country"`
	CustomerID uuid.UUID `json:"-"`
	Base
}

// BeforeSave will generate a UUID rather than numeric ID.
func (b *Base) BeforeSave(tx *gorm.DB) (err error) {
	b.UUID = uuid.New()
	return
}

// BeforeCreate will genenerate an AddressID
func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.AddressID = uuid.New()
	return
}

// BeforeCreate will generate a CustomerID uuid
func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.CustomerID = uuid.New()
	return
}
