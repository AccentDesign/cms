package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// AllowHead allows HEAD requests to be made against valid routes without
// the need to explicitly register a handler for each one.
func AllowHead() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodHead {
				c.Request().Method = http.MethodGet

				defer func() {
					c.Request().Method = http.MethodHead
				}()

				if err := next(c); err != nil {
					if err.Error() == echo.ErrMethodNotAllowed.Error() {
						return c.NoContent(http.StatusOK)
					}

					return err
				}
			}

			return next(c)
		}
	}
}
