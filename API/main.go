// Version: 1.0
package main

// Import required packages

// main - Entry point of the API
func main() {

	ConnectMongoDB()

	// Initialize SuperTokens
	SuperTokensInit()

	// Call the request handler function
	HandleRequests()
}
