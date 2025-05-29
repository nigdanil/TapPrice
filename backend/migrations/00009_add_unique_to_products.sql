-- +goose Up
ALTER TABLE products
ADD CONSTRAINT unique_product_name UNIQUE(name);

-- +goose Down
ALTER TABLE products
DROP CONSTRAINT unique_product_name;
