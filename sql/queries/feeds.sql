-- name: CreateFeed :one

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

SELECT * FROM feeds
WHERE url = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT *, u.name as username
FROM
    feeds f
LEFT JOIN
    users u ON f.user_id = u.id;