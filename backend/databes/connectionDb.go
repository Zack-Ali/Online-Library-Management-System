package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func ConnectDB() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file: ", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("Error from MongoDB Connection")
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error from Ping MongoDB Connection")
	}
	fmt.Println("Connection established successfully")
	database := client.Database("golang_online_books")
	Db = database
	return Db
}
