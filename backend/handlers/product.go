package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nigdanil/tapprice/models"
)

type ProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Composition string   `json:"composition"`
	Category    string   `json:"category"`
	Venue       string   `json:"venue"`
	CertLinks   []string `json:"cert_links"`
}

func CreateProductHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err := models.UpsertProduct(db, req.Name, req.Description, req.Composition, req.Category, req.Venue, req.CertLinks)
		if err != nil {
			http.Error(w, "Failed to create product", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func CreateProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqs []ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqs); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		for _, req := range reqs {
			err := models.UpsertProduct(db, req.Name, req.Description, req.Composition, req.Category, req.Venue, req.CertLinks)
			if err != nil {
				http.Error(w, "Failed to insert one of the products", http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("✅ Products created"))
	}
}

func DeleteProductHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err = models.DeleteProductByID(db, id)
		if err != nil {
			http.Error(w, "Failed to delete product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

type DeleteProductsRequest struct {
	IDs []int64 `json:"ids"`
}

func DeleteProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DeleteProductsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.IDs) == 0 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := models.DeleteProductsByID(db, req.IDs); err != nil {
			http.Error(w, "Failed to delete products", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ Products deleted"))
	}
}
