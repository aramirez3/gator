// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one

INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id)
    values ( $1, $2, $3, $4, $5, $6)
    RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const deleteAllFeeds = `-- name: DeleteAllFeeds :exec
DELETE FROM feeds
`

func (q *Queries) DeleteAllFeeds(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllFeeds)
	return err
}

const getFeed = `-- name: GetFeed :one

SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds
WHERE url = $1
`

func (q *Queries) GetFeed(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT f.id, f.created_at, f.updated_at, f.name, url, user_id, last_fetched_at, u.id, u.created_at, u.updated_at, u.name, u.name as username
FROM
    feeds f
LEFT JOIN
    users u ON f.user_id = u.id
`

type GetFeedsRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt sql.NullTime
	ID_2          uuid.NullUUID
	CreatedAt_2   sql.NullTime
	UpdatedAt_2   sql.NullTime
	Name_2        sql.NullString
	Username      sql.NullString
}

func (q *Queries) GetFeeds(ctx context.Context) ([]GetFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsRow
	for rows.Next() {
		var i GetFeedsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchedAt,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name_2,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds
    SET last_fetched_at = $2
    WHERE id = $1
`

type MarkFeedFetchedParams struct {
	ID            uuid.UUID
	LastFetchedAt sql.NullTime
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.ID, arg.LastFetchedAt)
	return err
}
