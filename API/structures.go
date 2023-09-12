// Version: 1.0
package main

// Data structure for an error json response
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

// Data structure for a user collection from MongoDB
type DBUser struct {
	ID             string `bson:"_id"`             //The user's unique ID provided by MongoDB and SuperTokens
	Handle         string `bson:"handle"`          //The user's handle
	CreatedAt      int64  `bson:"created_at"`      //Unix timestamp of when the user was created
	AccVerified    bool   `bson:"acc_verified"`    //Whether or not the user has a verified account
	EmailVerified  bool   `bson:"email_verified"`  //Whether or not the user has a verified email
	FollowerCount  int    `bson:"follower_count"`  //The number of followers the user has
	FollowingCount int    `bson:"following_count"` //The number of users the user is following
	LikeCount      int    `bson:"like_count"`      //The number of likes this user has on their posts
	RepostCount    int    `bson:"repost_count"`    //The number of reposts this user has on their posts
}

// Data structure for a user stats collection from MongoDB
type DBUserStats struct {
	ID        string   `bson:"_id"`       //The user's unique ID provided by MongoDB and SuperTokens
	Following []string `bson:"following"` //A slice of user IDs that the user is following
	Followers []string `bson:"followers"` //A slice of user IDs that are following the user
	Likes     []string `bson:"likes"`     //A slice of post IDs that the user has liked
	Reposts   []string `bson:"reposts"`   //A slice of post IDs that the user has reposted
}

// Data structure for a user content collection from MongoDB
type DBUserContent struct {
	ID    string   `bson:"_id"`   //The user's unique ID provided by MongoDB and SuperTokens
	Posts []string `bson:"posts"` //A slice of post IDs that the user has posted
}

// Data structure for a user json response
type UserResponse struct {
	ID              string `json:"user_id"`         //The user's unique ID provided by MongoDB and SuperTokens
	Handle          string `json:"handle"`          //The user's handle
	CreatedAt       int64  `json:"created_at"`      //Unix timestamp of when the user was created
	AccountVerified bool   `json:"acc_verified"`    //Whether or not the user has a verified account
	EmailVerified   bool   `json:"email_verified"`  //Whether or not the user has a verified email
	FollowerCount   int    `json:"follower_count"`  //The number of followers the user has
	FollowingCount  int    `json:"following_count"` //The number of users the user is following
	LikeCount       int    `json:"like_count"`      //The number of likes this user has on their posts
	RepostCount     int    `json:"repost_count"`    //The number of reposts this user has on their posts
}

// Data structure for a user content json response
type UserContentResponse struct {
	ID    string   `json:"user_id"` //The user's unique ID provided by MongoDB and SuperTokens
	Posts []string `json:"posts"`   //A slice of post IDs that the user has posted
}

// Data structure for a user stats json response
type UserStatsResponse struct {
	ID        string   `json:"user_id"`   //The user's unique ID provided by MongoDB and SuperTokens
	Following []string `json:"following"` //A slice of user IDs that the user is following
	Followers []string `json:"followers"` //A slice of user IDs that are following the user
	Likes     []string `json:"likes"`     //A slice of post IDs that the user has liked
	Reposts   []string `json:"reposts"`   //A slice of post IDs that the user has reposted
}

// Data structure for a signup request
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
