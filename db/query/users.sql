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

-- name: UpdateUserName :exec
UPDATE users
SET name = $2
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
