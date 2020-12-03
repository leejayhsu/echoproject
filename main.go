package main

import (
	berrors "echoproject/errors"
	"echoproject/handlers"
	"echoproject/middleware"
	"echoproject/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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

	// a custom error handler
	e.HTTPErrorHandler = berrors.CustomHTTPErrorHandler

	// register middleware
	e.Use(middleware.Auth)

	// set minimum log level (default is error)
	e.Logger.SetLevel(log.DEBUG)

	e.POST("/customers", handlers.CreateCustomer)
	e.GET("/customers/:customerID", handlers.GetCustomer)
	e.PATCH("/customers/:customerID", handlers.UpdateCustomer)
	e.DELETE("/customers/:customerID", handlers.DeleteCustomer)

	// for debugging and stuff
	e.POST("/e", handlers.TestError)
	e.POST("/e2", handlers.TestError2)

	var a string
	if os.Getenv("B_ENV") == "local" {
		// this is really just for mac's to avoid
		// annoying prompt to allow network connection for go app
		a = "localhost:8000"
	} else {
		a = ":8000"
	}
	e.Logger.Fatal(e.Start(a))
}
