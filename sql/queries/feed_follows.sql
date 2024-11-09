-- name: CreateFeedFollow :one

WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (
        id,
        created_at,
        updated_at,
        user_id,
        feed_id)
    VALUES ( 
        $1,
        $2,
        $3,
        $4,
        $5
        )
        RETURNING *
    )
SELECT
    inserted_feed_follows.*,
    f.name as feed_name,
    u.name as user_name
FROM inserted_feed_follows
INNER JOIN
    users u on inserted_feed_follows.user_id = u.id
INNER JOIN
    feeds f on inserted_feed_follows.feed_id = f.id;

-- name: GetFeedFollowsForUser :many

SELECT ff.*, f.name as feed_name 
FROM feed_follows ff
INNER JOIN
    feeds f on ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollowsForUserAndFeed :exec
DELETE FROM feed_follows
WHERE 
    user_id = $1 AND
    feed_id = $2;

-- name: DeleteAllFeedFollows :exec
DELETE FROM feed_follows;