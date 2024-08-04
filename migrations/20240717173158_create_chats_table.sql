-- +goose Up
-- +goose StatementBegin
CREATE TABLE Chats
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
CREATE TABLE Messages
(
    id         SERIAL PRIMARY KEY,
    chat_id    INT REFERENCES Chats (id) ON DELETE CASCADE,
    user_id    INT       NOT NULL,
    text       TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
CREATE TABLE ChatContributors
(
    chat_id    INT REFERENCES Chats (id) ON DELETE CASCADE,
    user_id    INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (chat_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Chats;
DROP TABLE Messages;
DROP TABLE ChatContributors;
-- +goose StatementEnd
