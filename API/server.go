// Version: 1.0
package main

// Import required packages
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"

	"github.com/joho/godotenv"
)

// main - Entry point of the API
func main() {

	// load .env file
	godotenv.Load(".env")

	// Initialize SuperTokens
	SuperTokensInit()

	// Call the request handler function
	HandleRequests()
}

// goDotEnvVariable - Loads .env file and returns the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// SuperTokensInit - Called by main() initializes SuperTokens
func SuperTokensInit() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// These are the connection details of the app you created on supertokens.com
			ConnectionURI: "https://st-dev-542a1560-4cc0-11ee-9fe8-cf09491f6e40.aws.supertokens.io",
			APIKey:        "UnuW=a5b8hWCRS9yxzP2KYvyRE",
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

	// Adding handlers.CORS(options)(supertokens.Middleware(router)))
	http.ListenAndServe(":8081", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))
}

// Post struct
type Post struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	Likes     int    `json:"likes"`
	Reposts   int    `json:"reposts"`
	Timestamp string `json:"timestamp"`
}

// Posts array
type Posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {

	posts := Posts{
		Post{
			Author:    "ibxcodecat",
			Content:   "I like cats!",
			Likes:     0,
			Reposts:   0,
			Timestamp: "2021-01-01 00:00:00",
		},
		{
			Author:    "ibxcodecat",
			Content:   "I like dogs!",
			Likes:     0,
			Reposts:   0,
			Timestamp: "2021-01-01 00:00:00",
		},
	}
	fmt.Println("Endpoint Hit: getPosts")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint Hit: createPost")
}

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OPEX API Hit")
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
