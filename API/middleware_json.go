// Version 1.0
package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// JSONMiddleware is a middleware function to process JSON requests
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request's Content-Type is application/json
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		// Create a struct to unmarshal the JSON data into
		var requestData map[string]interface{}

		// Decode the JSON data from the request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
			return
		}

		// Set the decoded JSON data in the request context
		ctx := context.WithValue(r.Context(), "json_data", requestData)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
