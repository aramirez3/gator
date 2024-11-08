-- name: CreateFeeds :one

INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id)
    values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
    )
    RETURNING *;

-- name: GetFeed :one

SELECT * from feeds
WHERE name = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT * from feeds;