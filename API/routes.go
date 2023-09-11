// Version: 1.0
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	var userReq UserRequest
	var dbUser DBUser

	err := DecodeJSONBody(w, r, &userReq)

	// Check for malformed request error and handle it
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if userReq.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "User ID not provided",
		})
		return
	}

	// Fetch the user from MongoDB
	dbUser, err = DBFetchUser(userReq.ID)

	fmt.Println(dbUser)

	// If MongoDB returns an empty DBUser check for errors and respond
	if dbUser.ID == "" {
		if err != nil {
			//There was a database error so therefore the database is unavailible
			fmt.Println("Error fetching user from MongoDB:", err)

			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(ErrorResponse{StatusCode: http.StatusServiceUnavailable, ErrorMessage: "User Database Unavailible"})
			return
		} else {
			//There was not a database error so the user was not found
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{StatusCode: http.StatusNotFound, ErrorMessage: "User not found"})
			return
		}
	}

	returnedAndProcessedUser := UserResponse{
		ID:              dbUser.ID,
		AccountVerified: dbUser.AccVerified,
		EmailVerified:   dbUser.EmailVerified,
		Handle:          dbUser.Handle,
		CreatedAt:       dbUser.CreatedAt,
	}

	// Send the response
	json.NewEncoder(w).Encode(returnedAndProcessedUser)
}
