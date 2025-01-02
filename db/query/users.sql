-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING user_id, name, created_at;

-- name: GetUserByID :one
SELECT user_id, name, created_at
FROM users
WHERE user_id = $1;

-- name: ListUsers :many
SELECT user_id, name, created_at
FROM users
ORDER BY created_at DESC;
