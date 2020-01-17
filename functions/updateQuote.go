package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rhyuen/golang_mongo_faas/model"
	"github.com/rhyuen/golang_mongo_faas/mw"
)

type ExpectedPayload struct {
	Text string `json:"text"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	id := mw.GetURLParams(r)

	var toUpdate ExpectedPayload
	json.NewDecoder(r.Body).Decode(&toUpdate)

	//nextText := toUpdate.Text

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	currQuote := model.Quote{Id: objId, Text: toUpdate.Text}
	err = currQuote.UpdateQuote(col)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Fprint(w, updateResult.ModifiedCount)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}
