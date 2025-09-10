-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cooking_durations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    delete_at TIMESTAMP
);

INSERT INTO
    cooking_durations (name, create_at, update_at)
VALUES
    ('5 - 10 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('10 - 30 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('30 - 60 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('More than 1 hour', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cooking_durations;
-- +goose StatementEnd
