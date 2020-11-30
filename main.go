package main

import (
	"echoproject/handlers"
	"echoproject/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=disable TimeZone=America/Los_Angeles"
	var err error
	models.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	models.DB.AutoMigrate(&models.Customer{})

	e := echo.New()
	e.POST("/customers", handlers.CreateCustomer)
	e.GET("/customers/:custID", handlers.GetCustomer)
	e.PATCH("/customers/:custID", handlers.UpdateCustomer)
	e.DELETE("/customers/:custID", handlers.DeleteCustomer)
	e.Logger.Fatal(e.Start("localhost:5000"))
}
