-- +goose Up
CREATE TABLE IF NOT EXISTS venues (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT UNIQUE,
    branding JSONB
);

-- +goose Down
DROP TABLE IF EXISTS venues;
