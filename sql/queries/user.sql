-- name: CreateUser :one
INSERT INTO users(id, name, api_key)
VALUES ($1, $2, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: UpdateUser :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users ORDER BY created_at DESC;