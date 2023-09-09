// Version: 1.0
package main

import (
	"encoding/json"
	"fmt"
)

// ProcessPostObjects - Processes post objects from MongoDB into a slice of PostResponse objects
func ProcessPostObjects(postObj []byte, f_err error) []PostResponse {

	// Check for errors while fetching posts
	if f_err != nil {
		fmt.Println("Error fetching posts from MongoDB:", f_err)
		return nil
	}

	// JSON data as a byte slice
	jsonData := []byte(postObj)

	// Create a slice of Post to store the result
	var posts []PostResponse

	// Unmarshal the JSON data into the posts slice
	err := json.Unmarshal(jsonData, &posts)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	// Create a slice of PostResponse to store the result
	var structuredPosts []PostResponse

	// Iterate over the posts
	for _, post := range posts {

		// Create a PostResponse object
		res := PostResponse{
			ID:         post.ID,
			Author:     post.Author,
			Content:    post.Content,
			Likes:      post.Likes,
			Reposts:    post.Reposts,
			RepostFrom: post.RepostFrom,
			Timestamp:  post.Timestamp,
		}

		// Append the post to the structuredPosts slice
		structuredPosts = append(structuredPosts, res)
	}

	// Print the result to the console for us to debug
	for _, post := range structuredPosts {
		fmt.Printf("\n\nPost:\n")
		fmt.Printf("ID: %s\n", post.ID)
		fmt.Printf("Author: %s\n", post.Author)
		fmt.Printf("Content: %s\n", post.Content)
		fmt.Printf("Likes: %d\n", post.Likes)
		fmt.Printf("RepostFrom: %v\n", post.RepostFrom)
		fmt.Printf("Reposts: %d\n", post.Reposts)
		fmt.Printf("Timestamp: %d\n", post.Timestamp)
	}

	// Return the result
	return structuredPosts

}

// ProcessUserObjects - Processes user objects from MongoDB into a slice of UserResponse objects
func ProcessUserObjects(userObj []byte, f_err error) []UserResponse {
	return nil
}
