package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection
var ctx context.Context

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27019")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the database and collection variables
	db = client.Database("blog").Collection("posts")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
}
