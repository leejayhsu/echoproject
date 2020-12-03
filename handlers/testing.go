package handlers

import (
	berrors "echoproject/errors"
	"errors"

	"github.com/labstack/echo"
)

// TestError is an endpoint for inducing a generic error
func TestError(c echo.Context) error {
	err := errors.New("asdf")
	return err
}

// TestError2 is an endpoint for inducing a Berror error
func TestError2(c echo.Context) error {
	extra := map[string](string){"status": "this is an extra"}
	err := &berrors.BError{
		Message: "this is a BError",
		Code:    424,
		Status:  "custom error",
		Kind:    "server error",
		Extra:   extra,
	}
	return err
}
