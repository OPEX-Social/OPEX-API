// Version: 1.0
package main

// Import required packages
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable - Loads .env file and returns the value of the key
func GoDotEnvVariable(key string) string {

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
