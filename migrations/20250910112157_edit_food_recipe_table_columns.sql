-- +goose Up
-- +goose StatementBegin
ALTER TABLE food_recipes
    RENAME COLUMN create_at TO created_at;

ALTER TABLE food_recipes
    RENAME COLUMN update_at TO updated_at;

ALTER TABLE food_recipes
    RENAME COLUMN delete_at TO deleted_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE food_recipes
    RENAME COLUMN created_at TO create_at;

ALTER TABLE food_recipes
    RENAME COLUMN updated_at TO update_at;

ALTER TABLE food_recipes
    RENAME COLUMN deleted_at TO delete_at;
-- +goose StatementEnd
