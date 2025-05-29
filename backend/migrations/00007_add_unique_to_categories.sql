-- +goose Up
ALTER TABLE categories
ADD CONSTRAINT unique_category_name UNIQUE(name);

-- +goose Down
ALTER TABLE categories
DROP CONSTRAINT unique_category_name;
