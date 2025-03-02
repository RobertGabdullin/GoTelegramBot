-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    chat_id INTEGER PRIMARY KEY
);
-- +goose StatementEnd