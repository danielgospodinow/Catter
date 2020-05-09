package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

// InitClient initializes the MongoDB client.
func InitClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

// GetClient retrieves the MongoDB client.
// Note: First call InitClient, in order to initialize the client.
func GetClient() *mongo.Client {
	return dbClient
}

// CloseClient closes the connection to MongoDB.
func CloseClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = dbClient.Disconnect(ctx)
}
