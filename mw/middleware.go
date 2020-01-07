package mw

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() (*mongo.Client, error) {
	url := os.Getenv("go_mongo_db")
	clientOptions := options.Client().ApplyURI(url).SetRetryWrites(false)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("InitDB connected to MongoDB.")

	return client, nil
}

func DBConnCollection(name, collection string) (*mongo.Collection, *mongo.Client, error) {
	url := os.Getenv("go_mongo_db")
	clientOptions := options.Client().ApplyURI(url).SetRetryWrites(false)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	fmt.Println("InitDB connected to MongoDB.")

	col := client.Database(name).Collection(collection)
	return col, client, nil
}

func GetURLParams(r *http.Request) string {
	return strings.Split(r.URL.Path, "/")[3]
}
