package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//models para serem usados exclusivamente pelo streaming do mongo
type DbEvent struct {
	DocumentKey   documentKey `bson:"documentKey"`
	OperationType string      `bson:"operationType"`
}

type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}
