package middleware

import (
	"net/http"
)

// MethodNotAllowedHandler is a middleware to handle 405 Method Not Allowed
func MethodNotAllowedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
