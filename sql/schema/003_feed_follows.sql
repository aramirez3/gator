-- +goose Up
CREATE table feed_follows(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL
        REFERENCES users (id)
        ON DELETE CASCADE,
    feed_id UUID NOT NULL
        REFERENCES feeds (id)
        ON DELETE CASCADE,
    UNIQUE (user_id, feed_id),
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (feed_id)
        REFERENCES feeds (id)
);

-- +goose Down
DROP table feed_follows;