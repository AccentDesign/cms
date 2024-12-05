// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: settings.sql

package dbx

import (
	"context"
)

const getSettings = `-- name: GetSettings :one
SELECT id, meta_description, meta_og_site_name, meta_og_title, meta_og_description, meta_og_url, meta_og_type, meta_og_image, meta_og_image_secure_url, meta_og_image_width, meta_og_image_height, meta_article_publisher, meta_article_section, meta_article_tag, meta_twitter_card, meta_twitter_image, meta_twitter_site, meta_robots, created_at, updated_at
FROM settings
LIMIT 1
`

// get the settings
func (q *Queries) GetSettings(ctx context.Context) (Setting, error) {
	row := q.db.QueryRow(ctx, getSettings)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.MetaDescription,
		&i.MetaOgSiteName,
		&i.MetaOgTitle,
		&i.MetaOgDescription,
		&i.MetaOgUrl,
		&i.MetaOgType,
		&i.MetaOgImage,
		&i.MetaOgImageSecureUrl,
		&i.MetaOgImageWidth,
		&i.MetaOgImageHeight,
		&i.MetaArticlePublisher,
		&i.MetaArticleSection,
		&i.MetaArticleTag,
		&i.MetaTwitterCard,
		&i.MetaTwitterImage,
		&i.MetaTwitterSite,
		&i.MetaRobots,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
