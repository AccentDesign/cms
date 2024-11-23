-- name: GetPageAncestors :many
-- get ancestors of a page
SELECT *
FROM page
WHERE path @> $1::ltree
ORDER BY level;

-- name: GetPageByPath :one
-- get a page by its path
SELECT sqlc.embed(p), tableoid::regclass::varchar as source
FROM page p
WHERE path = $1::ltree
LIMIT 1;

-- name: GetPageChildren :many
-- get the children of a page
SELECT *
FROM page
WHERE path <@ $1::ltree
  AND level = nlevel($1::ltree) + 1
ORDER BY path;

-- name: GetPageParent :one
-- get the parent of a page
SELECT *
FROM page
WHERE path = CASE WHEN nlevel($1::ltree) > 0 THEN subpath($1::ltree, 0, nlevel($1::ltree) - 1) END
LIMIT 1;

-- name: GetPageSiblings :many
-- get the siblings of a page
SELECT *
FROM page
WHERE path <@ subpath($1::ltree, 0, nlevel($1::ltree) - 1)
  AND level = nlevel($1::ltree)
  AND path <> $1::ltree
ORDER BY path;

-- name: GetPagesForSitemap :many
-- get the page for the sitemap
SELECT DISTINCT url, updated_at
FROM page
WHERE is_in_sitemap
ORDER BY url;
