package middleware

import (
	"fmt"
	"net/http"
)

func middlewareSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Security 1")
		next.ServeHTTP(w, r)
	})
}
