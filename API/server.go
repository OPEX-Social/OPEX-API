// Version: 1.0
package main

// Import required packages
import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// main - Entry point of the API
func main() {

	ConnectMongoDB()

	// Initialize SuperTokens
	SuperTokensInit()

	// Call the request handler function
	HandleRequests()

}

// SuperTokensInit - Called by main() initializes SuperTokens
func SuperTokensInit() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// These are the connection details of the app you created on supertokens.com
			ConnectionURI: GoDotEnvVariable("SUPERTOKENS_CONNECTION_URI"),
			APIKey:        GoDotEnvVariable("SUPERTOKENS_API_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "OPEX",
			APIDomain:       "http://localhost:8081",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
			dashboard.Init(nil),
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired, // or evmodels.ModeOptional
			}),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}

// HandleRequests - Called by main() handles all API requests
func HandleRequests() {
	// creates a new instance of a mux router and assigns it to router
	router := mux.NewRouter().StrictSlash(true)

	//Add our API routes and specify their respective functions and methods
	router.HandleFunc("/", Page)
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")

	router.HandleFunc("/user", GetUserByID).Methods("GET")

	// Adding handlers.CORS(options)(supertokens.Middleware(router))
	http.ListenAndServe(":8081", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))
}
