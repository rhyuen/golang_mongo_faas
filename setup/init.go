package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rhyuen/golang_mongo_faas/mw"
	"github.com/rhyuen/golang_mongo_faas/types"
)

func main() {
	client, err := mw.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("go_tester_one").Collection("quotes")

	err = collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collection Dropped from DB.")

	cap := types.Quote{"Captain America", "I can do this all day."}
	ghandi := types.Quote{"Mahatma Ghandi", "Be the change you wish to see in the world."}
	batman := types.Quote{"Batman", "vengeance is the night."}
	hulk := types.Quote{"Hulk", "I am always angry."}
	thor := types.Quote{"Thor", "More Mead!"}
	iron := types.Quote{"Iron Man", "I am Iron man."}
	spider := types.Quote{"Spider-Man", "Did you guys see that really old movie with the walkie thingies?"}

	dataList := make([]types.Quote, 0)
	dataList = append(dataList, cap, ghandi, batman, hulk, thor, iron, spider)

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
