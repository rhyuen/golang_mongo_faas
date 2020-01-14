package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rhyuen/golang_mongo_faas/model"
	"github.com/rhyuen/golang_mongo_faas/mw"
)

func main() {
	col, client, err := mw.DBConnCollection("go_tester_one", "quotes")
	if err != nil {
		log.Fatal(err)
	}

	err = col.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collection Dropped from DB.")

	file, err := ioutil.ReadFile("setup/data.json")
	if err != nil {
		fmt.Println("issue with reading file.")
		fmt.Println(err)
	}

	type File struct {
		Data []model.Quote `json:"data"`
	}
	var latest File
	err = json.Unmarshal(file, &latest)
	if err != nil {
		fmt.Println(err)
	}

	for _, q := range latest.Data {
		err := q.CreateQuote(col)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a document.")
	}

	client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InitDB Connection to MONGODB closed.")
}

func initWithFileData() error {
	return nil
}
