package pages

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"echo.go.dev/pkg/ui/pages"
)

// mapRelations maps dbx.Page relations to pages.Relation.
func mapRelations(relations []dbx.Page, settings dbx.Setting) []pages.Relation {
	mapped := make([]pages.Relation, len(relations))
	for i, rel := range relations {
		mapped[i] = pages.Relation{
			ID:        rel.ID,
			Title:     rel.Title,
			Path:      rel.Path,
			Url:       rel.Url.String,
			Level:     rel.Level.Int32,
			CreatedAt: rel.CreatedAt.Time,
			UpdatedAt: rel.UpdatedAt.Time,
			Meta:      getMeta(rel, settings),
		}
	}
	return mapped
}
