-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: InsertItem :one
INSERT INTO items (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllItems :many
SELECT * FROM items
ORDER BY name;