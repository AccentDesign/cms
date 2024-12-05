package pages

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"echo.go.dev/pkg/ui/pages"
	"github.com/jackc/pgx/v5/pgtype"
)

func getMeta(page dbx.Page, setting dbx.Setting) pages.Meta {
	return pages.Meta{
		Description:      getMetaField(page.MetaDescription, setting.MetaDescription),
		OGSiteName:       getMetaField(page.MetaOgSiteName, setting.MetaOgSiteName),
		OGTitle:          getMetaField(page.MetaOgTitle, setting.MetaOgTitle),
		OGDescription:    getMetaField(page.MetaOgDescription, setting.MetaOgDescription),
		OGUrl:            getMetaField(page.MetaOgUrl, setting.MetaOgUrl),
		OGType:           getMetaField(page.MetaOgType, setting.MetaOgType),
		OGImage:          getMetaField(page.MetaOgImage, setting.MetaOgImage),
		OGImageSecureUrl: getMetaField(page.MetaOgImageSecureUrl, setting.MetaOgImageSecureUrl),
		OGImageWidth:     getMetaField(page.MetaOgImageWidth, setting.MetaOgImageWidth),
		OGImageHeight:    getMetaField(page.MetaOgImageHeight, setting.MetaOgImageHeight),
		ArticlePublisher: getMetaField(page.MetaArticlePublisher, setting.MetaArticlePublisher),
		ArticleSection:   getMetaField(page.MetaArticleSection, setting.MetaArticleSection),
		ArticleTag:       getMetaField(page.MetaArticleTag, setting.MetaArticleTag),
		TwitterCard:      getMetaField(page.MetaTwitterCard, setting.MetaTwitterCard),
		TwitterImage:     getMetaField(page.MetaTwitterImage, setting.MetaTwitterImage),
		TwitterSite:      getMetaField(page.MetaTwitterSite, setting.MetaTwitterSite),
		Robots:           getMetaField(page.MetaRobots, setting.MetaRobots),
	}
}

func getMetaField(pageValue, settingValue pgtype.Text) string {
	if pageValue.Valid {
		return pageValue.String
	}
	return settingValue.String
}
