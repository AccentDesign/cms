package pages

import (
	"github.com/a-h/templ"
	"time"
)

// Meta items for a page.
type Meta struct {
	Description      string
	OGSiteName       string
	OGTitle          string
	OGDescription    string
	OGUrl            string
	OGType           string
	OGImage          string
	OGImageSecureUrl string
	OGImageWidth     string
	OGImageHeight    string
	ArticlePublisher string
	ArticleSection   string
	ArticleTag       string
	TwitterCard      string
	TwitterImage     string
	TwitterSite      string
}

// Relation is a relation of a page
type Relation struct {
	ID        int32
	Title     string
	Path      string
	Url       string
	Level     int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Meta      Meta
}

// PageType is an interface used as the page type Body is used to render the body tag.
type PageType interface {
	Body(p *Page) templ.Component
}

// Page struct.
type Page struct {
	ID        int32
	Title     string
	Path      string
	Url       string
	Level     int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Meta      Meta
	PageType  PageType
	Ancestors []Relation
	Children  []Relation
}

// PageTypeListing is used when the content listed from the children and only relevant to base pages.
type PageTypeListing struct{}

// PageTypeHTML is used when the content is derived from the table page_html
type PageTypeHTML struct {
	Html string
}
