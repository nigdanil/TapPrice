package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/nigdanil/tapprice/db"
	"github.com/nigdanil/tapprice/handlers"
	"github.com/nigdanil/tapprice/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env not loaded, using default env")
	}
}

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("❌ DB connection failed: %v", err)
	}
	defer database.Close()
	log.Println("✅ Connected to database")

	if err := db.Migrate(database); err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}
	log.Println("✅ Migrations applied")

	// 🧩 Создание первого администратора, если пользователей нет
	ensureInitialAdmin(database)

	r := mux.NewRouter()
	handlers.RegisterRoutes(r, database)

	// Apply CORS middleware
	handlerWithMiddleware := middleware.AuditLogger(database)(middleware.WithCORS(r))

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithMiddleware))
}

func ensureInitialAdmin(db *sql.DB) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	if err != nil {
		log.Fatalf("❌ Failed to check user count: %v", err)
	}

	if count == 0 {
		username := os.Getenv("INIT_ADMIN_USERNAME")
		password := os.Getenv("INIT_ADMIN_PASSWORD")
		role := os.Getenv("INIT_ADMIN_ROLE")

		if username == "" || password == "" || role == "" {
			log.Fatal("❌ INIT_ADMIN_* env variables not set")
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		_, err := db.Exec(`INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`,
			username, hash, role)
		if err != nil {
			log.Fatalf("❌ Failed to create initial admin: %v", err)
		}

		log.Printf("✅ Initial admin created: %s (role: %s)\n", username, role)
	}
}
