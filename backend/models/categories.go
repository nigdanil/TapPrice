package models

import (
	"database/sql"
	"fmt"
)

type Category struct {
	ID      int64  `json:"id"`
	VenueID *int64 `json:"venue_id,omitempty"` // указатель, потому что venue_id может быть NULL
	Name    string `json:"name"`
}

func GetAllCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query(`SELECT id, venue_id, name FROM categories`)
	if err != nil {
		return nil, fmt.Errorf("get categories: %w", err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		var venueID sql.NullInt64

		if err := rows.Scan(&c.ID, &venueID, &c.Name); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		if venueID.Valid {
			c.VenueID = &venueID.Int64
		}

		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategoryByID(db *sql.DB, id int64) (*Category, error) {
	var c Category
	var venueID sql.NullInt64

	row := db.QueryRow(`SELECT id, venue_id, name FROM categories WHERE id = $1`, id)
	err := row.Scan(&c.ID, &venueID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get category by id: %w", err)
	}
	if venueID.Valid {
		c.VenueID = &venueID.Int64
	}

	return &c, nil
}
func DeleteCategoryByID(db *sql.DB, id int64) error {
	result, err := db.Exec(`DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete category: %w", err)
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

func UpdateCategoryByID(db *sql.DB, id int64, c *Category) error {
	_, err := db.Exec(`UPDATE categories SET name = $1, venue_id = $2 WHERE id = $3`, c.Name, c.VenueID, id)
	if err != nil {
		return fmt.Errorf("update category: %w", err)
	}
	return nil
}
