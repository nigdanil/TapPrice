-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    venue_id INT REFERENCES venues(id),
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS categories;
