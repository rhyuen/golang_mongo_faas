package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rhyuen/golang_mongo_faas/model"
	"github.com/rhyuen/golang_mongo_faas/mw"
)

type ExpectedRequest struct {
	Id string `json:"id"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	id := mw.GetURLParams(r)

	//fmt.Println(id)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
		return
	}

	currQuote := model.Quote{Id: objectId}
	err = currQuote.DeleteQuote(col)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
		return
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
		return
		// errPayload := types.ErrorPayload{"deleteQuote Path", "Deletion Error", }
		// json.NewEncoder(w).Encode(errPayload)

	}

	fmt.Println("Delete Quote route disconnected.")
}
