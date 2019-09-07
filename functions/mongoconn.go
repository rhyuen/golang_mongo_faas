package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type Payloads struct {
	Path string    `json:"path"`
	Data []Trainer `json:"data"`
}

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("go_mongo_db")

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("connected to db.")
	collection := client.Database("go_tester_one").Collection("trainers")

	var firstOne Trainer

	filter := bson.D{{"age", 10}}
	collection.FindOne(context.TODO(), filter).Decode(&firstOne)

	list := make([]Trainer, 0)
	list = append(list, firstOne)
	send := Payloads{"/mongoconn", list}

	json.NewEncoder(w).Encode(send)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Closed.")
}
