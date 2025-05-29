-- +goose Up
CREATE TABLE audit_log (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL DEFAULT now(),
    ip TEXT,
    user_id INTEGER,
    path TEXT
);

-- +goose Down
DROP TABLE IF EXISTS audit_log;
