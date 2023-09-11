// Version: 1.0
package main

type ErrorResponse struct {
	StatusCode   int    `json:"status"`        //The HTTP status code
	ErrorMessage string `json:"error_message"` //The error message
}

// Data structure for a post json response
type PostResponse struct {
	ID         string `json:"_id"`         //The post's unique ID provided by MongoDB
	Author     string `json:"author"`      //A user ID string for the post author
	Content    string `json:"content"`     //The plaintext content of the post
	Likes      int    `json:"likes"`       //The number of likes the post has
	Reposts    int    `json:"reposts"`     //The number of reposts the post has
	RepostFrom string `json:"repost_from"` //If this is a repost, the user ID of the original post author
	Timestamp  int64  `json:"timestamp"`   //Unix timestamp of when the post was created
}

// Data structure for a user json request
type UserRequest struct {
	ID string `json:"user_id"`
}

// Data structure for a user object from MongoDB
type DBUser struct {
	ID            string `bson:"_id"`            //The user's unique ID provided by MongoDB and SuperTokens
	Handle        string `bson:"handle"`         //The user's handle
	CreatedAt     int64  `bson:"created_at"`     //Unix timestamp of when the user was created
	AccVerified   bool   `bson:"acc_verified"`   //Whether or not the user has a verified account
	EmailVerified bool   `bson:"email_verified"` //Whether or not the user has a verified email
}

// Data structure for a user json response
type UserResponse struct {
	ID              string `json:"_id"`            //The user's unique ID provided by MongoDB and SuperTokens
	Handle          string `json:"handle"`         //The user's handle
	CreatedAt       int64  `json:"created_at"`     //Unix timestamp of when the user was created
	AccountVerified bool   `json:"acc_verified"`   //Whether or not the user has a verified account
	EmailVerified   bool   `json:"email_verified"` //Whether or not the user has a verified email
}

// Data structure for a signup request
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
