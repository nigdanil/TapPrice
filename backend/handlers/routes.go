package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/nigdanil/tapprice/middleware"
	"github.com/nigdanil/tapprice/models"
)

// RegisterRoutes настраивает все маршруты API
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	// Public
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	r.Handle("/register", middleware.RequireAuth(http.HandlerFunc(RegisterHandler(db)))).Methods("POST")
	r.HandleFunc("/login", LoginHandler(db)).Methods("POST")

	r.HandleFunc("/products", getAllProductsHandler(db)).Methods("GET")
	r.HandleFunc("/product/{id}", getProductHandler(db)).Methods("GET")

	r.HandleFunc("/category/{id}", getCategoryHandler(db)).Methods("GET")
	r.Handle("/category/{id}", middleware.RequireAuth(DeleteCategoryHandler(db))).Methods("DELETE")

	r.HandleFunc("/categories", getAllCategoriesHandler(db)).Methods("GET")
	r.HandleFunc("/venue/{id}", getVenueHandler(db)).Methods("GET")
	r.Handle("/venue/{id}", middleware.RequireAuth(DeleteVenueHandler(db))).Methods("DELETE")
	r.HandleFunc("/venues", getAllVenuesHandler(db)).Methods("GET")

	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	// Защищённый маршрут — только для admin
	r.Handle("/admin-only", middleware.RequireRole("admin", http.HandlerFunc(AdminOnlyHandler))).Methods("GET")

	r.Handle("/users", middleware.RequireRole("admin", GetAllUsersHandler(db))).Methods("GET")
	r.Handle("/user/{id}", middleware.RequireRole("admin", DeleteUserHandler(db))).Methods("DELETE")
	r.Handle("/user/{id}", middleware.RequireRole("admin", UpdateUserHandler(db))).Methods("PUT")

	r.Handle("/change-password", middleware.RequireAuth(ChangePasswordHandler(db))).Methods("POST")

	r.Handle("/products/import", middleware.RequireAuth(http.HandlerFunc(ImportProductsHandler(db)))).Methods("POST")
	r.Handle("/product", middleware.RequireAuth(http.HandlerFunc(CreateProductHandler(db)))).Methods("POST")
	r.Handle("/product/{id}", middleware.RequireAuth(http.HandlerFunc(DeleteProductHandler(db)))).Methods("DELETE")
	r.Handle("/products/delete", middleware.RequireAuth(http.HandlerFunc(DeleteProductsHandler(db)))).Methods("POST")

	r.Handle("/products", middleware.RequireAuth(http.HandlerFunc(CreateProductsHandler(db)))).Methods("POST")
	r.Handle("/audit-log", middleware.RequireRole("admin", GetAuditLogHandler(db))).Methods("GET")

	r.Handle("/category/{id}", middleware.RequireAuth(UpdateCategoryHandler(db))).Methods("PUT")
	r.Handle("/venue/{id}", middleware.RequireAuth(UpdateVenueHandler(db))).Methods("PUT")
	r.Handle("/product/{id}", middleware.RequireAuth(UpdateProductHandler(db))).Methods("PUT")

}

func getProductHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		product, err := models.GetProductByID(db, id)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func getAllProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var venueID, categoryID *int64

		if v := r.URL.Query().Get("venue_id"); v != "" {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				http.Error(w, "Invalid venue_id", http.StatusBadRequest)
				return
			}
			venueID = &id
		}

		if c := r.URL.Query().Get("category_id"); c != "" {
			id, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				http.Error(w, "Invalid category_id", http.StatusBadRequest)
				return
			}
			categoryID = &id
		}

		products, err := models.GetAllProducts(db, venueID, categoryID)
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func getAllCategoriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📥 Запрос: GET /categories")

		categories, err := models.GetAllCategories(db)
		if err != nil {
			log.Printf("❌ Ошибка при получении категорий: %v", err)
			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	}
}

func getAllVenuesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📥 Запрос: GET /venues")

		venues, err := models.GetAllVenues(db)
		if err != nil {
			log.Printf("❌ Ошибка при получении площадок: %v", err)
			http.Error(w, "Failed to fetch venues", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(venues)
	}
}

func getCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		category, err := models.GetCategoryByID(db, id)
		if err != nil {
			http.Error(w, "Failed to get category", http.StatusInternalServerError)
			return
		}
		if category == nil {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(category)
	}
}

func getVenueHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		venue, err := models.GetVenueByID(db, id)
		if err != nil {
			http.Error(w, "Failed to get venue", http.StatusInternalServerError)
			return
		}
		if venue == nil {
			http.Error(w, "Venue not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(venue)
	}
}

func DeleteCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		err = models.DeleteCategoryByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Category not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to delete category", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Category deleted successfully"}`))
	}
}

func DeleteVenueHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid venue ID", http.StatusBadRequest)
			return
		}

		err = models.DeleteVenueByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Venue not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to delete venue", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Venue deleted successfully"}`))
	}
}

func UpdateCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		var cat models.Category
		if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := models.UpdateCategoryByID(db, id, &cat); err != nil {
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Category updated successfully"}`))
	}
}

func UpdateVenueHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid venue ID", http.StatusBadRequest)
			return
		}

		var venue models.Venue
		if err := json.NewDecoder(r.Body).Decode(&venue); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := models.UpdateVenueByID(db, id, &venue); err != nil {
			http.Error(w, "Failed to update venue", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Venue updated successfully"}`))
	}
}

func UpdateProductHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := models.UpdateProductByID(db, id, &product); err != nil {
			http.Error(w, "Failed to update product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Product updated successfully"}`))
	}
}
