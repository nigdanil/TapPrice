-- +goose Up
CREATE TABLE IF NOT EXISTS prices (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    price NUMERIC(10,2) NOT NULL,
    currency TEXT DEFAULT '₽',
    updated_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS prices;
