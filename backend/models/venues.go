package models

import (
	"database/sql"
	"fmt"
)

type Venue struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug,omitempty"`
}

func GetAllVenues(db *sql.DB) ([]Venue, error) {
	rows, err := db.Query(`SELECT id, name, slug FROM venues`)
	if err != nil {
		return nil, fmt.Errorf("get venues: %w", err)
	}
	defer rows.Close()

	var venues []Venue
	for rows.Next() {
		var v Venue
		var slug sql.NullString

		if err := rows.Scan(&v.ID, &v.Name, &slug); err != nil {
			return nil, fmt.Errorf("scan venue: %w", err)
		}
		if slug.Valid {
			v.Slug = slug.String
		}

		venues = append(venues, v)
	}

	return venues, nil
}
