// Version: 1.0
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Posts array
type Posts []PostResponse

func getPosts(w http.ResponseWriter, r *http.Request) {

	/*
		postsSampleData := Posts{
			PostResponse{
				Author:    "ibxcodecat",
				Content:   "I like cats!",
				Likes:     0,
				Reposts:   0,
				Timestamp: 200,
			},
			{
				Author:    "ibxcodecat",
				Content:   "I like dogs!",
				Likes:     0,
				Reposts:   0,
				Timestamp: 200,
			},
		}
	*/

	fmt.Println("Endpoint Hit: getPosts")

	// Fetch all posts from MongoDB
	db_posts, err := DBFetchAllPosts(GoDotEnvVariable("MONGO_DB_NAME"), GoDotEnvVariable("MONGO_POST_COLLECTION"))
	if err != nil {
		fmt.Println("Error fetching posts from MongoDB:", err)
		return
	}

	// Process the post objects from MongoDB into a slice of PostResponse objects
	posts := ProcessPostObjects(json.Marshal(db_posts))

	// Send the response
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint Hit: createPost")
}

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OPEX API Hit")
}
