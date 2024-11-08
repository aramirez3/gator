-- name: CreateFeedFollow :one

INSERT INTO feed_follows (
    id,
    created_at,
    updated_at,
    user_id,
    feed_id)
    values (
    $1,
    $2,
    $3,
    $4,
    $5
    )
    RETURNING *;

-- name: GetFeedFollowsFowUser :one

SELECT * FROM feed_follows
WHERE user_id = $1;

-- name: DeleteAllFeedFollowss :exec
DELETE FROM feed_follows;