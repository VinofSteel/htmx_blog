-- name: ListAllUsers :many
SELECT * FROM users ORDER BY $1 OFFSET $2 LIMIT $3;