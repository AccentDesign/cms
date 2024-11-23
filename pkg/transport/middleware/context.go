package middleware

import (
	"echo.go.dev/pkg/config"
	"echo.go.dev/pkg/storage/db/dbx"
	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Config   *config.Config
	HTMX     *HTMX
	Postgres *pgxpool.Pool
	Queries  *dbx.Queries
}

func (c *CustomContext) RenderComponent(statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	return c.HTML(statusCode, buf.String())
}

// Context middleware func to define a custom context.
func Context(postgres *pgxpool.Pool, config *config.Config) echo.MiddlewareFunc {
	queries := dbx.New(postgres)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			htmx := &HTMX{Request: c.Request(), Response: c.Response()}
			cc := &CustomContext{
				Context:  c,
				Config:   config,
				HTMX:     htmx,
				Postgres: postgres,
				Queries:  queries,
			}
			return next(cc)
		}
	}
}
