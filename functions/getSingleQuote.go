package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rhyuen/golang_mongo_faas/mw"
	"github.com/rhyuen/golang_mongo_faas/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("go_tester_one").Collection("quotes")

	fmt.Println(strings.Split(r.URL.Path, "/")[3])
	id := mw.GetURLParams(r)

	fmt.Println(id)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objId}}

	var getResult types.Quote
	err = collection.FindOne(context.TODO(), filter).Decode(&getResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The number of items retrieved: ", getResult)

	data := make([]types.Quote, 0)
	data = append(data, getResult)

	payload := types.Payload{"/getSingleQuote", data}
	json.NewEncoder(w).Encode(payload)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
