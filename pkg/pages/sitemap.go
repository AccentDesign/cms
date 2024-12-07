package pages

import (
	"context"
	"echo.go.dev/pkg/storage/db/dbx"
	"encoding/xml"
	"strings"
	"time"
)

type SitemapEntry struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	Priority   float32   `xml:"priority,omitempty"`
	ChangeFreq string    `xml:"changefreq,omitempty"`
}

type Sitemap struct {
	XMLName   xml.Name        `xml:"urlset"`
	Namespace string          `xml:"xmlns,attr"`
	Entries   []*SitemapEntry `xml:"url"`
}

func getSitemap(ctx context.Context, queries *dbx.Queries) (*Sitemap, error) {
	sitemap := &Sitemap{
		Namespace: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Entries:   []*SitemapEntry{},
	}

	entries, err := queries.GetPagesForSitemap(ctx)
	if err != nil {
		return sitemap, err
	}

	for _, entry := range entries {
		sitemap.Entries = append(sitemap.Entries, &SitemapEntry{
			Loc:        strings.TrimSuffix(entry.Url.String, "/"),
			LastMod:    entry.UpdatedAt.Time,
			Priority:   entry.Priority,
			ChangeFreq: string(entry.ChangeFrequency),
		})
	}

	return sitemap, nil
}
