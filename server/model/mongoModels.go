package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//models to be used exclusively on mongoDB.
type DbEvent struct {
	DocumentKey   documentKey `bson:"documentKey"`
	OperationType string      `bson:"operationType"`
}

type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}
