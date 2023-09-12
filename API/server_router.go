// Version: 1.0
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// HandleRequests - Called by main() handles all API requests
func HandleRequests() {
	// creates a new instance of a mux router and assigns it to router
	router := mux.NewRouter().StrictSlash(true)

	//Add our API routes and specify their respective functions and methods
	router.HandleFunc("/", Page)

	rate_limiter := NewRateLimiter()

	// Create a new APIRouteInfo struct for each route

	/*
		route_getPosts := APIRouteInfo{
			AllowedRequestMethods: []string{METHOD_GET},
			LimiterWindow:         time.Second * 5,
			LimiterMax:            2,
			Path:                  "/posts",
		}*/

	// Create the get user route @ api/user - GET
	route_getMyUser := APIRouteInfo{
		AllowedRequestMethods: []string{METHOD_GET},
		LimiterWindow:         time.Second * 5,
		LimiterMax:            2,
		Path:                  "/users/@me",
	}

	// Create the get user route @ api/user - GET
	route_getAnyUser := APIRouteInfo{
		AllowedRequestMethods: []string{METHOD_GET},
		LimiterWindow:         time.Second * 5,
		LimiterMax:            2,
		Path:                  "/users/{id}",
	}

	// Create the create user route @ api/user - POST
	route_createNewUser := APIRouteInfo{
		AllowedRequestMethods: []string{METHOD_POST},
		LimiterWindow:         time.Second * 120,
		LimiterMax:            2,
		Path:                  "/users",
	}

	// Create the update user route @ api/user - PUT
	route_updateUser := APIRouteInfo{
		AllowedRequestMethods: []string{METHOD_PUT},
		LimiterWindow:         time.Second * 120,
		LimiterMax:            20,
		Path:                  "/users/@me",
	}

	// Create the delete user route @ api/user - DELETE
	route_deleteUser := APIRouteInfo{
		AllowedRequestMethods: []string{METHOD_DELETE},
		LimiterWindow:         time.Second * 10,
		LimiterMax:            2,
		Path:                  "/users/@me",
	}

	// Add the get user route to the router
	router.Handle(
		route_getMyUser.Path,
		MethodNotAllowedMiddleware(route_getMyUser.AllowedRequestMethods)(
			rate_limiter.RateLimit(
				route_getMyUser.Path,
				route_getMyUser.LimiterMax,
				route_getMyUser.LimiterWindow.Abs(),
				session.VerifySession(
					nil,
					HandleGetAuthenticatedUser,
				),
			),
		),
	)

	// Add the get user route to the router
	router.Handle(
		route_getAnyUser.Path,
		MethodNotAllowedMiddleware(route_getAnyUser.AllowedRequestMethods)(
			ParseFormMiddleware(
				rate_limiter.RateLimit(
					route_getAnyUser.Path,
					route_getAnyUser.LimiterMax,
					route_getAnyUser.LimiterWindow.Abs(),
					http.HandlerFunc(GetUserByID),
				),
			),
		),
	)

	// Add the create user route to the router
	router.Handle(
		route_createNewUser.Path,
		MethodNotAllowedMiddleware(route_createNewUser.AllowedRequestMethods)(
			rate_limiter.RateLimit(
				route_createNewUser.Path,
				route_createNewUser.LimiterMax,
				route_createNewUser.LimiterWindow.Abs(),
				session.VerifySession(
					nil,
					CreateNewUser,
				),
			),
		),
	)

	// Add the update user route to the router
	router.Handle(
		route_updateUser.Path,
		MethodNotAllowedMiddleware(route_updateUser.AllowedRequestMethods)(
			rate_limiter.RateLimit(
				route_updateUser.Path,
				route_updateUser.LimiterMax,
				route_updateUser.LimiterWindow.Abs(),
				session.VerifySession(
					nil,
					UpdateUser,
				),
			),
		),
	)

	// Add the delete user route to the router
	router.Handle(
		route_deleteUser.Path,
		MethodNotAllowedMiddleware(route_deleteUser.AllowedRequestMethods)(
			rate_limiter.RateLimit(
				route_deleteUser.Path,
				route_deleteUser.LimiterMax,
				route_deleteUser.LimiterWindow.Abs(),
				session.VerifySession(
					nil,
					DeleteUser,
				),
			),
		),
	)

	// Adding handlers.CORS(options)(supertokens.Middleware(router))
	http.ListenAndServe(":8081", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{
			METHOD_GET,
			METHOD_POST,
			METHOD_PUT,
			METHOD_DELETE,
		}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))
}

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OPEX API Hit")
}
