-- name: GetPageTypeHtml :one
-- get the data for a html page type
SELECT *
FROM page_html
WHERE id = $1
LIMIT 1;
