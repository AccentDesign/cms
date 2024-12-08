package static

import (
	"embed"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed public/*
	Public embed.FS
)

// Router create a new static router.
func Router(e *echo.Echo) {
	e.StaticFS("/static", echo.MustSubFS(Public, "public"))
}
