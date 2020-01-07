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

	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	var requestBody model.Quote
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("issue with decoding.")
		log.Fatal(err)
	}

	err = requestBody.CreateQuote(col)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateQuote Connection to MONGODB closed.")
}
