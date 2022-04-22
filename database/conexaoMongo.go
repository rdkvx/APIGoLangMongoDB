package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTIONSTRING = "mongodb://root:123mudar@127.0.0.1:27017/?authSource=admin"
	DB               = "desafiotecnico"
	COLLECTION       = "moedacripto"
)

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("FAILED TO CLOSE CONNECTION: ", err)
		}
	}()
}

func Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		fmt.Println(err)
	}

	return client, ctx, cancel, err
}

func GetCollection() (*mongo.Collection, error) {
	client, ctx, cancel, err := Connect()
	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database(DB).Collection(COLLECTION)

	defer Close(client, ctx, cancel)

	return collection, nil
}
