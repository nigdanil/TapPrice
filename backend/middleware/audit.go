package middleware

import (
	"database/sql"
	"log"
	"net"
	"net/http"
)

func AuditLogger(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var userID *int // nullable для неавторизованных
			if val := r.Context().Value(ContextUserID); val != nil {
				if uid, ok := val.(int); ok {
					userID = &uid
				}
			}

			host, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				host = r.RemoteAddr // fallback
			}

			_, err = db.Exec(`
				INSERT INTO audit_log (ip, user_id, path)
				VALUES ($1, $2, $3)
			`, host, userID, r.URL.Path)

			if err != nil {
				log.Printf("⚠️ Failed to write audit log: %v", err)
			}

			next.ServeHTTP(w, r)
		})
	}
}
