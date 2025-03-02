-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
    link_id SERIAL PRIMARY KEY,
    link TEXT NOT NULL,
    last_update TIMESTAMP
)
-- +goose StatementEnd