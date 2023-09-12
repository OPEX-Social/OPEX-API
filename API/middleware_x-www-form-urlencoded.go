// Version 1.0
package main

import "net/http"

func ParseFormMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request's Content-Type is application/x-www-form-urlencoded
		contentType := r.Header.Get("Content-Type")

		// If the Content-Type is not application/x-www-form-urlencoded return an error
		if contentType != "application/x-www-form-urlencoded" {
			http.Error(w, "Content-Type must be application/x-www-form-urlencoded", http.StatusUnsupportedMediaType)
			return
		}

		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
