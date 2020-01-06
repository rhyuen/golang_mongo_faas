package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rhyuen/golang_mongo_faas/mw"
	"go.mongodb.org/mongo-driver/bson"
)

type ExpectedPayload struct {
	Text string `json:"text"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("go_tester_one").Collection("quotes")

	fmt.Println(strings.Split(r.URL.Path, "/")[3])
	id := mw.GetURLParams(r)

	var toUpdate ExpectedPayload
	json.NewDecoder(r.Body).Decode(&toUpdate)

	nextText := toUpdate.Text

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objId}}
	update := bson.D{
		{"$set", bson.D{
			{"text", nextText},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The number of items update: ", updateResult.ModifiedCount)

	fmt.Fprint(w, updateResult.ModifiedCount)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}
