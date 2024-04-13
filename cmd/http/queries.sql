-- items
-- name: CreateItem :one
INSERT INTO items (
  id, site, price, smart_time, name, description, nickname
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;
