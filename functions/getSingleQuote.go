package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhyuen/golang_mongo_faas/model"
	"github.com/rhyuen/golang_mongo_faas/mw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Body struct {
	Path string        `json:"path"`
	Data []model.Quote `json:"quotes"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	id := mw.GetURLParams(r)

	fmt.Println(id)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	toUpdate := model.Quote{Id: objId}
	err = toUpdate.GetQuote(col)
	if err != nil {
		fmt.Println("issue with GetQuote")
		log.Fatal(err)

	}
	data := make([]model.Quote, 0)
	data = append(data, toUpdate)

	payload := Body{"/getSingleQuoteUpdated", data}
	json.NewEncoder(w).Encode(payload)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
