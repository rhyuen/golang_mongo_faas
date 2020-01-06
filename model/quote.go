package model

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Quote struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (q *Quote) GetQuote(db *mongo.Collection) error {
	
	return errors.New("not Implemented")
}

func (q *Quote) UpdateQuote(db *mongo.Collection) error {
	return errors.New("not implemeneted")
}

func (q *Quote) DeleteQuote(db *mongo.Collection) error {

	return errors.New("Not done")
}

func (q *Quote) CreateQuote(db *mongo.Collection) error {
	insertResult, err := db.InsertOne(context.TODO(), q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a document.", insertResult.InsertedID)

	return nil
}

func getQuotes(db *mongo.Collection) ([]Quote, error) {
	return nil, errors.New("not done yet.")
}
