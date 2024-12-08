package pages

import (
	"github.com/a-h/templ"
	"time"
)

// Meta items for a page.
type Meta struct {
	Description      string `json:"description"`
	OGSiteName       string `json:"og_site_name"`
	OGTitle          string `json:"og_title"`
	OGDescription    string `json:"og_description"`
	OGUrl            string `json:"og_url"`
	OGType           string `json:"og_type"`
	OGImage          string `json:"og_image"`
	OGImageSecureUrl string `json:"og_image_secure_url"`
	OGImageWidth     string `json:"og_image_width"`
	OGImageHeight    string `json:"og_image_height"`
	ArticlePublisher string `json:"article_publisher"`
	ArticleSection   string `json:"article_section"`
	ArticleTag       string `json:"article_tag"`
	TwitterCard      string `json:"twitter_card"`
	TwitterImage     string `json:"twitter_image"`
	TwitterSite      string `json:"twitter_site"`
	Robots           string `json:"robots"`
}

// Relation is a relation of a page
type Relation struct {
	ID            int32     `json:"id"`
	Title         string    `json:"title"`
	Path          string    `json:"path"`
	Url           string    `json:"url"`
	Level         int32     `json:"level"`
	FeaturedImage string    `json:"featured_image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Meta          Meta      `json:"meta"`
}

// PageType is an interface used as the page type Body is used to render the body tag.
type PageType interface {
	Body(page *Page) templ.Component
}

// Page struct.
type Page struct {
	ID            int32      `json:"id"`
	Title         string     `json:"title"`
	Path          string     `json:"path"`
	Url           string     `json:"url"`
	Level         int32      `json:"level"`
	Tags          []string   `json:"tags"`
	Categories    []string   `json:"categories"`
	FeaturedImage string     `json:"featured_image"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	PublishedAt   time.Time  `json:"published_at"`
	Meta          Meta       `json:"meta"`
	PageType      PageType   `json:"-"`
	Ancestors     []Relation `json:"ancestors"`
	Children      []Relation `json:"children"`
}

// PageTypeListing is used when the content listed from the children and only relevant to base pages.
type PageTypeListing struct{}

type SearchResult struct {
	ID              int32
	Title           string
	MetaDescription string
	Url             string
	Headline        string
	Rank            float32
}

// PageTypeSearch is used for the site search.
type PageTypeSearch struct {
	Query   string
	Results []SearchResult
}

// PageTypeHTML is used when the content is derived from the table page_html
type PageTypeHTML struct {
	Html string
}
