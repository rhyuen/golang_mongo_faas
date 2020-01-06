package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhyuen/golang_mongo_faas/mw"
	"github.com/rhyuen/golang_mongo_faas/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {

	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	var data []types.Quote

	collection := client.Database("go_tester_one").Collection("quotes")
	findOptions := options.Find()
	currItr, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for currItr.Next(context.TODO()) {
		var currQuote types.Quote

		err := currItr.Decode(&currQuote)

		if err != nil {
			log.Fatal(err)
		}
		data = append(data, currQuote)
	}

	if err := currItr.Err(); err != nil {
		log.Fatal(err)
	}

	currItr.Close(context.TODO())

	payload := types.Payload{"/getQuotes", data}
	json.NewEncoder(w).Encode(payload)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("getQuotes Connection to MONGODB closed.")

}
