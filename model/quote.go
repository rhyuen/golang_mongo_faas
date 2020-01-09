package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Quote struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author string             `json:"author"`
	Text   string             `json:"text"`
}

func (q *Quote) GetQuote(db *mongo.Collection) error {
	filter := bson.D{{"_id", q.Id}}
	err := db.FindOne(context.TODO(), filter).Decode(q)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (q *Quote) UpdateQuote(db *mongo.Collection) error {
	filter := bson.D{{"_id", q.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"text", q.Text},
		}},
	}

	updateResult, err := db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Issue with Update Quote")
		log.Fatal(err)
	}
	fmt.Println("The number of items update: ", updateResult.ModifiedCount)
	return nil
}

func (q *Quote) DeleteQuote(db *mongo.Collection) error {
	filter := bson.D{{"_id", q.Id}}
	deleteResult, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Issue with Delete Quote")
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the quotes collection.", deleteResult.DeletedCount)
	return nil
}

func (q *Quote) CreateQuote(db *mongo.Collection) error {
	insertResult, err := db.InsertOne(context.TODO(), q)
	if err != nil {
		fmt.Println("Issue with creating new Quote")
		log.Fatal(err)
	}
	fmt.Println("Inserted a document.", insertResult.InsertedID)

	return nil
}

func GetQuotes(db *mongo.Collection) ([]Quote, error) {
	findOptions := options.Find()
	currItr, err := db.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("issue with gettin all quotes.")
		log.Fatal(err)
	}

	var data []Quote
	for currItr.Next(context.TODO()) {
		var currQuote Quote

		err := currItr.Decode(&currQuote)

		if err != nil {
			fmt.Println("issue with iterating through quotes.")
			log.Fatal(err)
			return nil, err
		}
		data = append(data, currQuote)
	}

	if err := currItr.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	currItr.Close(context.TODO())

	return data, nil

}
