// Version: 1.0
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MethodNotAllowedMiddleware(allowedMethods []string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request's method is allowed
			methodAllowed := false
			for _, method := range allowedMethods {
				if r.Method == method {
					methodAllowed = true
					break
				}
			}

			if !methodAllowed {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
