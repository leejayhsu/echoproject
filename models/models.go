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
	// Addresses  []Address `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"addresses"`
	Addresses []Address `gorm:"foreignKey:CustomerID;references:CustomerID;" json:"addresses"`
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
	// log.Debug("In the customer constructor")
	log.Println("In the customer constructor")
	c.CustomerID = uuid.New()
	return
}

// BeforeDelete soft deletes addresses
func (c *Customer) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("In the customer delete hook")
	// it's just print zero values...
	log.Println(c)
	log.Println(c.FirstName)
	log.Println("done printing")
	// if result := tx.Model(&Address{}).Where("customer_id = ?", c.CustomerID).Update("street1", "deleted!"); result.Error != nil {
	// 	log.Println(result.Error)
	// }
	// tx.Model(&Customer{}).Where("customer_id = ?", c.CustomerID).Update("first_name", "this was modified by the delete hook")
	// c.FirstName = "deleted by hook"
	// DB.Model(&Address{}).Where("customer_id = ?", c.CustomerID).Update("street1", "deleted")

	return
}
