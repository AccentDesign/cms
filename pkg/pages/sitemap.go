package pages

import (
	"context"
	"echo.go.dev/pkg/config"
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

func getSitemap(ctx context.Context, queries *dbx.Queries, cfg *config.Config) (*Sitemap, error) {
	sitemap := &Sitemap{
		Namespace: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Entries:   []*SitemapEntry{},
	}

	entries, err := queries.GetPagesForSitemap(ctx)
	if err != nil {
		return sitemap, err
	}

	for _, entry := range entries {
		loc := strings.TrimSuffix(cfg.Server.Url, "/")
		if entry.Url.String != "/" {
			loc += entry.Url.String
		}
		sitemap.Entries = append(sitemap.Entries, &SitemapEntry{
			Loc:     loc,
			LastMod: entry.UpdatedAt.Time,
		})
	}

	return sitemap, nil
}
