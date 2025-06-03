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

func GetVenueByID(db *sql.DB, id int64) (*Venue, error) {
	var v Venue
	var slug sql.NullString

	row := db.QueryRow(`SELECT id, name, slug FROM venues WHERE id = $1`, id)
	err := row.Scan(&v.ID, &v.Name, &slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get venue by id: %w", err)
	}
	if slug.Valid {
		v.Slug = slug.String
	}

	return &v, nil
}
func DeleteVenueByID(db *sql.DB, id int64) error {
	result, err := db.Exec(`DELETE FROM venues WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete venue: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func UpdateVenueByID(db *sql.DB, id int64, v *Venue) error {
	_, err := db.Exec(`UPDATE venues SET name = $1, slug = $2 WHERE id = $3`, v.Name, v.Slug, id)
	if err != nil {
		return fmt.Errorf("update venue: %w", err)
	}
	return nil
}
