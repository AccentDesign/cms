package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS returns an Echo middleware function for cors.
func CORS(cfg config.SecurityConfig) echo.MiddlewareFunc {
	corsConfig := middleware.CORSConfig{
		AllowOrigins: cfg.AllowedHosts,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
		Skipper: func(c echo.Context) bool {
			return c.Path() != "/*"
		},
	}

	return middleware.CORSWithConfig(corsConfig)
}
