-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('admin', 'operator', 'manager')),
    created_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE users;
