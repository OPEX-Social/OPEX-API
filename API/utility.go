// Version: 1.0
package main

// Import required packages
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// HTTP Methods used in the API
const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

const (
	FORM_KEY_USER_ID = "user_id"
	FORM_KEY_HANDLE  = "handle"
)

// APIRouteInfo - Information about an API route
type APIRouteInfo struct {
	AllowedRequestMethods []string      // The request method for this route
	Path                  string        // The path for this route
	LimiterWindow         time.Duration // The fixed time window for the rate limit
	LimiterMax            int           // The number of requests allowed in the time window
}

// EnvVar - Loads .env file and returns the value of the key
func EnvVar(key string) string {

	fmt.Println("Reloading .env file")

	// Load .env file
	err := godotenv.Load("./.env")

	// If there is an error loading the .env file, log the error and exit
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Return the value of the key
	return os.Getenv(key)
}
