package pages

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/ui/pages"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"sync"
)

func Router(e *echo.Echo) {
	g := e.Group("")
	{
		g.GET("", pageHandler)
		g.GET("/sitemap.xml", sitemapHandler)
		g.GET("/*", pageHandler)
	}
}

func pageHandler(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := c.Request().Context()

	var (
		wg                                              sync.WaitGroup
		settings                                        dbx.Setting
		page                                            dbx.GetPageByPathRow
		ancestors, children                             []dbx.Page
		errSettings, errPage, errAncestors, errChildren error
	)

	path := normalizePath(c.Request().URL.Path)

	wg.Add(4)

	go func() {
		defer wg.Done()
		settings, errSettings = cc.Queries.GetSettings(ctx)
	}()

	go func() {
		defer wg.Done()
		page, errPage = cc.Queries.GetPageByPath(ctx, path)
	}()

	go func() {
		defer wg.Done()
		ancestors, errAncestors = cc.Queries.GetPageAncestors(ctx, path)
	}()

	go func() {
		defer wg.Done()
		children, errChildren = cc.Queries.GetPageChildren(ctx, path)
	}()

	wg.Wait()

	if errSettings != nil || errPage != nil || errAncestors != nil || errChildren != nil {
		return echo.NotFoundHandler(c)
	}

	method, ok := pageTypeFactory[page.Source]
	if !ok {
		return fmt.Errorf("unsupported page type: %s", page.Source)
	}

	pageType, err := method(c, cc.Queries, &page.Page)
	if err != nil {
		return fmt.Errorf("failed to fetch page type with ID %d: %w", page.Page.ID, err)
	}

	pageComponent := &pages.Page{
		ID:        page.Page.ID,
		Title:     page.Page.Title,
		Path:      page.Page.Path,
		Url:       page.Page.Url.String,
		Level:     page.Page.Level.Int32,
		CreatedAt: page.Page.CreatedAt.Time,
		UpdatedAt: page.Page.UpdatedAt.Time,
		Meta:      getMeta(page.Page, settings),
		PageType:  pageType,
		Ancestors: mapRelations(ancestors, settings),
		Children:  mapRelations(children, settings),
	}

	return cc.RenderComponent(http.StatusOK, pageComponent.HTML())
}

func normalizePath(rawPath string) string {
	path := strings.ToLower(rawPath)
	path = strings.Trim(path, "/")
	return strings.ReplaceAll(path, "/", ".")
}

func sitemapHandler(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	sitemap, err := getSitemap(c.Request().Context(), cc.Queries, cc.Config)
	if err != nil {
		c.Error(err)
	}

	return c.XML(http.StatusOK, sitemap)
}
