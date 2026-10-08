-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeedByName :one
SELECT * FROM feeds
WHERE name = $1 LIMIT 1;

-- name: GetUserFeeds :many
SELECT * FROM feeds
WHERE user_id = $1;
