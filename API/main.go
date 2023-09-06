package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OPEX API Hit")
}

func HandleRequests() {
	http.HandleFunc("/", Page)
	http.HandleFunc("/posts", getPosts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	HandleRequests()
}
