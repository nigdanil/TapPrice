package db

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB) error {
	return goose.Up(db, "./migrations")
}
