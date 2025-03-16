package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db = "go-mongo2"
const coll = "user"

func ConnectDB() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading env file")
	}
	MongoDB := os.Getenv("MONGODB_URL")

	conOption := options.Client().ApplyURI(MongoDB)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, conOption)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("database connection successful!")
	return client
}

var client = ConnectDB()

func GetCollection() *mongo.Collection {
	return client.Database(db).Collection(coll)

}
