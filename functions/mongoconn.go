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
)

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
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
