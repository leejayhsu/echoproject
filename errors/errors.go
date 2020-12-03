package errors

import (
	"net/http"

	"github.com/labstack/echo"
)

// BError implements Error
type BError struct {
	Message string
	Code    int
	Status  string
	Kind    string
	Extra   map[string](string) `json:"-"`
}

func (e *BError) Error() string {
	return e.Message
}

// CustomHTTPErrorHandler is an error handler for the service
func CustomHTTPErrorHandler(err error, c echo.Context) {
	type response struct {
		Message string
		Code    int
		Status  string
		Kind    string
		Extra   map[string](string) `json:"-"`
	}

	code := http.StatusInternalServerError

	// check if error is an echo HTTPError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message := he.Message.(string)
		r := response{Message: message, Code: 500, Status: "unknown error", Kind: "server error"}
		c.JSON(code, r)
	}

	if be, ok := err.(*BError); ok {
		c.JSON(be.Code, be)
	}
	c.Logger().Error(err)

}
