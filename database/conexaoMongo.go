package database

import (
	//"DesafioTecnico/server/model"

	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	//"gopkg.in/mgo.v2/bson"
)

func Uri()(conexao string){
	return "mongodb://root:123mudar@127.0.0.1:27017/?authSource=admin"
}

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

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}

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
	result, err := collection.InsertOne(context.Background(),doc)
	if err != nil {
		fmt.Print("Erro ao inserir: ",err)
	}
	
	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the database and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	
	return
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, filter, option interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on filter and option.
	result, err = collection.Find(context.Background(), filter)
	return
}