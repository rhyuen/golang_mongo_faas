package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rhyuen/golang_mongo_faas/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("go_mongo_db")
	clientOptions := options.Client().ApplyURI(url).SetRetryWrites(false)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db.")

	var requestBody types.Quote
	json.NewDecoder(r.Body).Decode(&requestBody)

	collection := client.Database("go_tester_one").Collection("quotes")
	insertResult, err := collection.InsertOne(context.TODO(), requestBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a document.", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateQuote Connection to MONGODB closed.")
}
