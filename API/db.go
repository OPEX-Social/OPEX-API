// Version: 1.0
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() {

	// Set client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := GoDotEnvVariable("MONGO_URI")
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// REQUIRED FOR DB CONNECTION, NOBODY HAS ANY IDEA WHAT THIS DOES, BUT IT WORKS SO DON'T TOUCH IT OR ELSE!
	// This function was found on an obscure Russian 1995 developer forum by Nathan Schmitt (thank him later)
	// Next suffering dev increment this counter to warn future devs | HOURS DEBUGGING - 6.25
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged MongoDB Successfully. Connected to MongoDB!")
}
