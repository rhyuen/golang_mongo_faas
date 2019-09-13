package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rhyuen/golang_mongo_faas/mw"
	"go.mongodb.org/mongo-driver/bson"
)

type ExpectedRequest struct {
	Id string `json:"id"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("go_tester_one").Collection("quotes")
	//TODO: Change below to query param from form.

	var expectedId ExpectedRequest
	json.NewDecoder(r.Body).Decode(&expectedId)

	objectId, err := primitive.ObjectIDFromHex(expectedId.Id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
		return
	}

	filter := bson.D{{"_id", objectId}}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
		return
	}

	fmt.Printf("Deleted %v documents in the quotes collection.", deleteResult.DeletedCount)

	fmt.Fprintf(w, string(deleteResult.DeletedCount))

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
