// Version 1.0
package main

import "time"

// HTTP Methods used in the API
const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

// APIRouteInfo - Information about an API route
type APIRouteInfo struct {
	RequestMethod string        // The request method for this route
	Path          string        // The path for this route
	LimiterWindow time.Duration // The fixed time window for the rate limit
	LimiterMax    int           // The number of requests allowed in the time window
}
