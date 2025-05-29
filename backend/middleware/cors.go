package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func WithCORS(handler http.Handler) http.Handler {
	return cors.AllowAll().Handler(handler)
}
