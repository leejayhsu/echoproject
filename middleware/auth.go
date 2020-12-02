package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

// Auth is the middleware function.
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Logger().Info("in auth middleware")

		i := c.Request().Header.Get("Identity")
		a := c.Request().Header.Get("Authorization")

		if i != "A" {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthenticated")
		}
		if a != "A" {
			return echo.NewHTTPError(http.StatusForbidden, "unauthorized")
		}

		// if auth passed, call next
		return next(c)
	}
}
