# Adding new page types

For reference this is adding another like the `html` page type.

## 1. define your table to store the data in a new migration file

Migrations are in: `pkg/storage/db/migrations`.

```sql
create table page_article
(
    name varchar(100) not null,
    company varchar(100) not null
) inherits (page);

create trigger set_updated_at
    before update
    on page_article
    for each row
    execute procedure cms_set_updated_at();

create trigger check_path_uniqueness
    before insert or update
    on page_article
    for each row
    execute procedure cms_page_path_uniqueness();
```

Once done migrate the database to see your new table.

## 2. add your sql to query the data

Sql code is in: `pkg/storage/db/sqlc/page_type.sql`.

Based in the above an example would be:

```sql
-- name: GetPageTypeArticle :one
-- get the data for a article page type
SELECT *
FROM page_article
WHERE id = $1
LIMIT 1;
```

Then run the sqlc gen task:

```bash
task gen:sqlc
```

This will generate the golang code to query the data.

## 3. Add your new page type struct

Types defined in: `pkg/ui/pages/types.go`.

Example templ func in: `pkg/ui/pages/page_html.templ`.

Your struct will look like:

```go
type PageTypeArticle struct {
	Name string
	Company string
}
```

Your templ func to render the `Body` like:

```go
package pages

templ (t *PageTypeArticle) Body(p *Page) {
	<body class="grid gap-6 p-6 antialiased">
		<header class="space-y-6">
			<div class="space-y-1">
				<h1 class="owl-h1">{ p.Title }</h1>
				<div>{ p.Meta.Description }</div>
			</div>
			@p.Breadcrumb()
		</header>
		<main class="space-y-6">
			<div>{ t.Name }</div>
			<div>{ t.Company }</div>
		</main>
		@p.Footer()
	</body>
}
```

## 4. Add the factory function to get the page type for the page

Factory code in: `pkg/pages/factory.go`.

```go
var pageTypeFactory = map[string]func(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error){
    ...
    "page_article": handlePageArticle,
}

func handlePageArticle(c echo.Context, queries *dbx.Queries, page *dbx.Page) (pages.PageType, error) {
	row, err := queries.GetPageTypeArticle(c.Request().Context(), page.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page_article with ID %d: %w", page.ID, err)
	}
	return &pages.PageTypeArticle{
		Name: row.Name,
		Company: row.Company,
	}, nil
}
```

Any pages added to your `page_article` table should now load on their urls.
