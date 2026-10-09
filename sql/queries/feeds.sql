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

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: GetUserFeeds :many
SELECT feeds.* FROM feeds
JOIN users on feeds.user_id = users.id
WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * from feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at asc NULLS FIRST;
