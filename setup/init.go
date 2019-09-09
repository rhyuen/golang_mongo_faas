package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rhyuen/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	url := os.Getenv("go_mongo_db")
	clientOptions := options.Client().ApplyURI(url).SetRetryWrites(false)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("InitDB connected to MongoDB.")

	collection := client.Database("go_tester_one").Collection("quotes")

	cap := types.Quote{"Captain America", "I can do this all day."}
	ghandi := types.Quote{"Mahatma Ghandi", "Be the change you wish to see in the world."}
	batman := types.Quote{"Batman", "vengeance is the night."}

	dataList := make([]types.Quote, 0)
	dataList = append(dataList, cap, ghandi, batman)

	for _, q := range dataList {
		insertResult, err := collection.InsertOne(context.TODO(), q)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a document.", insertResult.InsertedID)
	}

	client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InitDB Connection to MONGODB closed.")
}
