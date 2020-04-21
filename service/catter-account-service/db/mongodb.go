package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var dbClient *mongo.Client

func InitClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

func GetClient() *mongo.Client {
	return dbClient
}

func CloseClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = dbClient.Disconnect(ctx)
}
