-- name: GetCSSClasses :many
-- returns a list of css classes found in the html content.
WITH extracted AS (
    SELECT unnest(regexp_matches(html, 'class="([^"]*)"', 'g')) AS classes
    FROM page_html
)
SELECT DISTINCT unnest(string_to_array(classes, ' ')) AS class
FROM extracted
ORDER BY class;