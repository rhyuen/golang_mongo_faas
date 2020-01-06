package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Path string  `json:"path"`
	Data []Quote `json:"quotes"`
}

type ErrorPayload struct {
	Path        string `json:"path"`
	Description string `json:"error"`
	Message     string `json:"details"`
}

type Quote struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author string             `json:"author" bson:"author"`
	Text   string             `json:"text" bson:"text"`
}
