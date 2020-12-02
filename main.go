package main

import (
	berrors "echoproject/errors"
	"echoproject/handlers"
	"echoproject/models"
	"errors"
	"net/http"
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

	// a custom error handler
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.POST("/customers", handlers.CreateCustomer)
	e.GET("/customers/:customerID", handlers.GetCustomer)
	e.PATCH("/customers/:customerID", handlers.UpdateCustomer)
	e.DELETE("/customers/:customerID", handlers.DeleteCustomer)

	e.POST("/e", testError)
	e.POST("/e2", testError2)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	type response struct {
		Message string
		Code    int
		Status  string
		Kind    string `json:"-"`
		Extra   map[string](string)
	}

	code := http.StatusInternalServerError
	// check if error is an echo HTTPError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		c.JSON(code, "bond error not raised")
	}

	if be, ok := err.(*berrors.BError); ok {
		c.JSON(be.Code, be)
	}
	c.Logger().Error(err)

}

func testError(c echo.Context) error {
	err := errors.New("asdf")
	return err
}

func testError2(c echo.Context) error {
	extra := map[string](string){"status": "this is an extra"}
	err := &berrors.BError{
		Message: "this is a bond error",
		Code:    424,
		Status:  "custom error",
		Kind:    "server error",
		Extra:   extra,
	}
	return err
}
