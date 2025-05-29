-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    venue_id INT REFERENCES venues(id),
    category_id INT REFERENCES categories(id),
    name TEXT NOT NULL,
    description TEXT,
    composition TEXT,
    cert_links TEXT[],
    updated_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS products;
    