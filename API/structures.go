// Version: 1.0
package main

// Data structure for a post json response
type PostResponse struct {
	Author     string `json:"author"`      //A user ID string for the post author
	Content    string `json:"content"`     //The plaintext content of the post
	Likes      int    `json:"likes"`       //The number of likes the post has
	Reposts    int    `json:"reposts"`     //The number of reposts the post has
	RepostFrom string `json:"repost_from"` //If this is a repost, the user ID of the original post author
	Timestamp  string `json:"timestamp"`   //Unix timestamp of when the post was created
}

// Data structure for a user json response
type UserResponse struct {
	UserID          string `json:"_id"`              //The user's unique ID provided by SuperTokens
	Handle          string `json:"handle"`           //The user's handle
	CreatedAt       string `json:"created_at"`       //Unix timestamp of when the user was created
	AccountVerified bool   `json:"account_verified"` //Whether or not the user has a verified account
	EmailVerified   bool   `json:"email_verified"`   //Whether or not the user has a verified email
}

// Data structure for a signup request
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
