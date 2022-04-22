package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//models to be used exclusively on mongoDB.
type DbEvent struct {
	DocumentKey   documentKey `json:"documentKey" bson:"documentKey"`
	OperationType string      `json:"operationType" bson:"operationType"`
}

type documentKey struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}
