package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

// Auth is the middleware function.
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Logger().Info("in auth middleware")

		k := c.Request().Header.Get("X-Api-Key")

		if k != "A" {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthenticated")
		}

		// if auth passed, call next
		return next(c)
	}
}
