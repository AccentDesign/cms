package pages

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"echo.go.dev/pkg/ui/pages"
	"fmt"
	"github.com/labstack/echo/v4"
)

// pageTypeFactory maps page table to their corresponding handler functions.
var pageTypeFactory = map[string]func(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error){
	"page":      handlePage,
	"page_html": handlePageHtml,
}

// handlePage get the page type for non db related content.
func handlePage(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	switch page.PageType {
	case dbx.PageTypeListing:
		return &pages.PageTypeListing{}, nil
	}
	return nil, fmt.Errorf("unsupported page type: %s", page.PageType)
}

// handlePageHtml get the page type for html.
func handlePageHtml(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	row, err := queries.GetPageTypeHtml(c.Request().Context(), page.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page_html with ID %d: %w", page.ID, err)
	}
	return &pages.PageTypeHTML{
		Html: row.Html,
	}, nil
}
