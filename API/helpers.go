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
