package pages

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"echo.go.dev/pkg/ui/pages"
	"fmt"
	"github.com/labstack/echo/v4"
)

// pageTypeFactory maps page table to their corresponding handler functions.
var pageTypeFactory = map[string]func(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error){
	"page":      handleGenericPage,
	"page_html": handlePageHtml,
}

// handleGenericPage get the page type for the base page type only.
func handleGenericPage(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	switch page.PageType {
	case dbx.PageTypeListing:
		return handleGenericPageListing(c, queries, page)
	case dbx.PageTypeSearch:
		return handleGenericPageSearch(c, queries, page)
	}
	return nil, fmt.Errorf("unsupported page type: %s", page.PageType)
}

// handleGenericPageListing get the generic listing page type.
func handleGenericPageListing(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	return &pages.PageTypeListing{}, nil
}

// handleGenericPageSearch get the generic search page type.
func handleGenericPageSearch(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	query := c.QueryParam("q")

	results, err := queries.GetPageSearchResults(c.Request().Context(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get search results: %w", err)
	}

	pageType := pages.PageTypeSearch{
		Results: make([]pages.SearchResult, len(results)),
		Query:   query,
	}

	for i, result := range results {
		pageType.Results[i] = pages.SearchResult{
			ID:              result.ID,
			Title:           result.Title,
			MetaDescription: result.MetaDescription.String,
			Url:             result.Url.String,
			Headline:        result.Headline,
			Rank:            result.Rank,
		}
	}

	return &pageType, nil
}

// handlePageHtml get the html page type.
func handlePageHtml(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	row, err := queries.GetPageTypeHtml(c.Request().Context(), page.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page_html with ID %d: %w", page.ID, err)
	}
	return &pages.PageTypeHTML{
		Html: row.Html,
	}, nil
}
