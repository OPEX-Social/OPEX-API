// Version: 1.0
package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoCLient - Set by ConnectToMongoDB()
var mongoClient *mongo.Client

// ConnectMongoDB - Connects to MongoDB
func ConnectMongoDB() {

	// Set client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := GoDotEnvVariable("MONGO_URI")
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	// Set the global mongoClient variable to the client
	mongoClient = client

	if err != nil {
		panic(err)
	}

	/*
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()*/

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged MongoDB Successfully. Connected to MongoDB!")
}

// ProcessPostObjects - Retrieves post objects from MongoDB
func DBFetchAllPosts(dbName, collectionName string) ([]bson.M, error) {
	// Get a handle to the "Posts" collection
	collection := mongoClient.Database(dbName).Collection(collectionName)

	// Find and fetch all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func DBFetchUser(requestedUserID string) (DBUser, error) {

	db_name := GoDotEnvVariable("MONGO_DB_NAME")
	collection_name := GoDotEnvVariable("MONGO_USER_COLLECTION")

	// Get a handle to the "Users" collection
	collection := mongoClient.Database(db_name).Collection(collection_name)

	filter := bson.M{"_id": requestedUserID}

	var dbUser DBUser

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find and fetch the document in the collection
	err := collection.FindOne(ctx, filter).Decode(&dbUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return DBUser{}, nil
		} else {
			return DBUser{}, err
		}
	}

	fmt.Println("Found a single document:", dbUser)

	// Construct and Return the user
	dbUser = DBUser{
		ID:             requestedUserID,
		AccVerified:    dbUser.AccVerified,
		EmailVerified:  dbUser.EmailVerified,
		Handle:         dbUser.Handle,
		CreatedAt:      dbUser.CreatedAt,
		FollowerCount:  dbUser.FollowerCount,
		FollowingCount: dbUser.FollowingCount,
		LikeCount:      dbUser.LikeCount,
		RepostCount:    dbUser.RepostCount,
	}

	return dbUser, nil
}
