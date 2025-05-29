package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/lib/pq"
)

type Product struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Composition string         `json:"composition"`
	CertLinks   pq.StringArray `json:"cert_links"`
}

// Получить один продукт по ID
func GetProductByID(db *sql.DB, id int64) (*Product, error) {
	query := `SELECT id, name, description, composition, cert_links FROM products WHERE id = $1`
	row := db.QueryRow(query, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Composition, &p.CertLinks)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return &p, nil
}

// Получить все продукты или отфильтрованные по venue_id
func GetAllProducts(db *sql.DB, venueID, categoryID *int64) ([]Product, error) {
	query := `SELECT id, name, description, composition, cert_links FROM products WHERE 1=1`
	args := []interface{}{}
	i := 1

	if venueID != nil {
		query += ` AND venue_id = $` + strconv.Itoa(i)
		args = append(args, *venueID)
		i++
	}
	if categoryID != nil {
		query += ` AND category_id = $` + strconv.Itoa(i)
		args = append(args, *categoryID)
		i++
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Composition, &p.CertLinks); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func UpsertProduct(db *sql.DB, name, desc, composition, categoryName, venueName string, certLinks []string) error {
	var catID, venueID int64

	err := db.QueryRow(`INSERT INTO categories (name) VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id`, categoryName).Scan(&catID)
	if err != nil {
		return err
	}

	err = db.QueryRow(`INSERT INTO venues (name) VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id`, venueName).Scan(&venueID)
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO products (name, description, composition, category_id, venue_id, cert_links)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (name) DO UPDATE SET
		description = EXCLUDED.description,
		composition = EXCLUDED.composition,
		category_id = EXCLUDED.category_id,
		venue_id = EXCLUDED.venue_id,
		cert_links = EXCLUDED.cert_links
	`, name, desc, composition, catID, venueID, pq.Array(certLinks))

	return err
}

func DeleteProductByID(db *sql.DB, id int64) error {
	_, err := db.Exec(`DELETE FROM products WHERE id = $1`, id)
	return err
}

func DeleteProductsByID(db *sql.DB, ids []int64) error {
	query := `DELETE FROM products WHERE id = ANY($1)`
	_, err := db.Exec(query, pq.Array(ids))
	return err
}
