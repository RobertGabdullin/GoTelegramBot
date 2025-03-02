-- +goose Up
-- +goose StatementBegin
CREATE TABLE chat_link (
    chat_link_id SERIAL PRIMARY KEY,
    chat_id INTEGER,
    link_id INTEGER,
    CONSTRAINT fk_chat_id FOREIGN KEY (chat_id) REFERENCES chat(chat_id) ON DELETE CASCADE,
    CONSTRAINT fk_link_id FOREIGN KEY (link_id) REFERENCES link(link_id) ON DELETE CASCADE
)
-- +goose StatementEnd
