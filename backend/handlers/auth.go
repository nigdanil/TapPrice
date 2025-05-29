package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nigdanil/tapprice/middleware"
	"github.com/nigdanil/tapprice/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

// список допустимых IP/подсетей, загружается из .env
var allowedCIDRs = strings.Split(os.Getenv("ALLOWED_IPS"), ",")

func isAllowedIPCIDR(remoteAddr string, cidrs []string) bool {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return false
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return false
	}

	for _, cidr := range cidrs {
		_, subnet, err := net.ParseCIDR(strings.TrimSpace(cidr))
		if err != nil {
			continue
		}
		if subnet.Contains(ip) {
			return true
		}
	}
	return false
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка IP
		fmt.Println("DEBUG RemoteAddr:", r.RemoteAddr)
		fmt.Println("DEBUG Allowed CIDRs:", allowedCIDRs)

		if !isAllowedIPCIDR(r.RemoteAddr, allowedCIDRs) {
			http.Error(w, "Forbidden: unauthorized IP", http.StatusForbidden)
			return
		}

		// Проверка роли из контекста
		role := r.Context().Value(middleware.ContextRole)
		if role != "admin" {
			http.Error(w, "Forbidden: admin only", http.StatusForbidden)
			return
		}

		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		_, err := db.Exec(`INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`,
			req.Username, hash, req.Role)
		if err != nil {
			http.Error(w, "User exists or error", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		var id int
		var hash, role string
		err := db.QueryRow(`SELECT id, password_hash, role FROM users WHERE username=$1`, req.Username).
			Scan(&id, &hash, &role)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, _ := middleware.GenerateJWT(id, role)

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute),
			SameSite: http.SameSiteStrictMode,
		})

		w.WriteHeader(http.StatusOK)
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
	w.WriteHeader(http.StatusOK)
}

func AdminOnlyHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middleware.ContextRole)
	if role != "admin" {
		http.Error(w, "Access denied: admin only", http.StatusForbidden)
		return
	}
	w.Write([]byte("🔐 Welcome, admin. You have full access."))
}

func ChangePasswordHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDVal := r.Context().Value(middleware.ContextUserID)
		if userIDVal == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userID := userIDVal.(int)

		var req struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var hash string
		err := db.QueryRow(`SELECT password_hash FROM users WHERE id = $1`, userID).Scan(&hash)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.OldPassword)) != nil {
			http.Error(w, "Invalid old password", http.StatusUnauthorized)
			return
		}

		newHash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err := models.UpdatePassword(db, userID, string(newHash)); err != nil {
			http.Error(w, "Failed to update password", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🔒 Password updated successfully"))
	}
}
