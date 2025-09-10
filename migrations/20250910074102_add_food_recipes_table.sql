-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS food_recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    ingredients TEXT NOT NULL,
    instructions TEXT NOT NULL,
    image_url TEXT NULL,
    difficulty_id INT NOT NULL REFERENCES difficulties(id),
    cooking_duration_id INT NOT NULL REFERENCES cooking_durations(id),
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    delete_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS food_recipes;
-- +goose StatementEnd
