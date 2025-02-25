package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client
var UserCollection *mongo.Collection

func ConnectMongoDB() {
	uri := "mongodb://torza:thnvaza123@host.docker.internal:27017/GolangCRUD?authSource=admin"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping Failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	DBClient = client
	UserCollection = client.Database("GolangCRUD").Collection("users")
}
