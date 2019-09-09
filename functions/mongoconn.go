package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rhyuen/golang_mongo_faas/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("go_mongo_db")

	clientOptions := options.Client().ApplyURI(url).SetRetryWrites(false)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db.")
	collection := client.Database("go_tester_one").Collection("quotes")

	var firstQuote types.Quote

	filter := bson.D{{"age", 10}}
	collection.FindOne(context.TODO(), filter).Decode(&firstQuote)

	list := make([]types.Quote, 0)
	list = append(list, firstQuote)
	send := types.Payload{"/mongoconn", list}

	json.NewEncoder(w).Encode(send)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Closed.")
}
