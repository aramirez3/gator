-- name: CreateFeed :one

INSERT INTO feed (
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

SELECT * from feed
WHERE name = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feed;