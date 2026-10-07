-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY name asc;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;
