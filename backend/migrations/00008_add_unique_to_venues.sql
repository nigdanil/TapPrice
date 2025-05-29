-- +goose Up
ALTER TABLE venues
ADD CONSTRAINT unique_venue_name UNIQUE(name);

-- +goose Down
ALTER TABLE venues
DROP CONSTRAINT unique_venue_name;
