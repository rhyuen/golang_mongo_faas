package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhyuen/golang_mongo_faas/model"
	"github.com/rhyuen/golang_mongo_faas/mw"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	var requestBody model.Quote
	json.NewDecoder(r.Body).Decode(&requestBody)

	// fmt.Println(requestBody.Author)
	// fmt.Println(requestBody.Text)

	// collection := client.Database("go_tester_one").Collection("quotes")
	// insertResult, err := collection.InsertOne(context.TODO(), requestBody)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a document.", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateQuote Connection to MONGODB closed.")
}
