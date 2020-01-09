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

type Body struct {
	Path string        `json:"path"`
	Data []model.Quote `json:"quotes"`
}

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {

	// client, err := mw.DBConnect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//collection := client.Database("go_tester_one").Collection("quotes")

	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	// var data []types.Quote

	// findOptions := options.Find()
	// currItr, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for currItr.Next(context.TODO()) {
	// 	var currQuote types.Quote

	// 	err := currItr.Decode(&currQuote)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	data = append(data, currQuote)
	// }

	// if err := currItr.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// currItr.Close(context.TODO())

	// payload := types.Payload{"/getQuotes", data}
	// json.NewEncoder(w).Encode(payload)

	data, err := model.GetQuotes(col)
	if err != nil {
		fmt.Println("issue with GetQuotes at Handler Level")
		log.Fatal(err)
	}

	payload := Body{"/getSingleQuoteUpdated", data}
	json.NewEncoder(w).Encode(payload)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("getQuotes Connection to MONGODB closed.")

}
