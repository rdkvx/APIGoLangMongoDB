package database

import (
	//"DesafioTecnico/server/model"

	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}


func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func ConnectMongo() {

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := Connect("mongodb+srv://rdkvx:Dragonforce123@cluster0.ntd1x.mongodb.net/desafiotecnico?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	Ping(client, ctx)
}
