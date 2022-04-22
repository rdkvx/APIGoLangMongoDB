package repositorio

import (
	"DesafioTecnico/database"
	m "DesafioTecnico/server/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func Create(mc m.CryptoCurrency) error {
	client, ctx, cancel, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	_, err = collection.InsertOne(context.Background(), mc)

	if err != nil {
		fmt.Print("FAILED TO INSERT: ", err.Error())
	}

	defer database.Close(client, ctx, cancel)

	return err
}

func Read(id string) (mc m.CryptoCurrency, err error) {
	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	//return a pointer that brings a object
	//corresponding to the id passed by parameters.
	cur := collection.FindOne(context.Background(), filter)
	if err != nil {
		return mc, err
	}

	var results = m.CryptoCurrency{}

	if err = cur.Decode(&results); err != nil {
		return results, err
	}

	return results, err
}

func ReadAll(sortParam string, ascending bool) (obj []m.CryptoCurrency, err error) {
	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	var cur *mongo.Cursor
	findOptions := options.Find()

	switch sortParam{
	case "id", "name", "symbol", "votes", "createdat", "updatedat": 
		if ascending { //If ascending = true, the return will be the default.
			findOptions.SetSort(bson.M{sortParam: 1})
			cur, err = collection.Find(context.Background(), bson.M{}, findOptions)
			if err != nil {
				return obj, err
			}
		} else { //Sort by `sortParam` field descending (higher value first)
			findOptions.SetSort(bson.M{sortParam: -1})
			cur, err = collection.Find(context.Background(), bson.M{}, findOptions)
			if err != nil {
				return obj, err
			}
		}
	default: 
		//Return a list on default order, wich is ordered by name asc
		sortParam = "name"
		findOptions.SetSort(bson.M{sortParam: 1})
		cur, err = collection.Find(context.Background(), bson.M{}, findOptions)
		if err != nil {
			return obj, err
		}
	}

	var results = []m.CryptoCurrency{}

	if err = cur.All(context.Background(), &results); err != nil {
		return obj, err
	}

	return results, err
}

func Update(mc m.CryptoCurrency) error {
	client, ctx, cancel, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	defer database.Close(client, ctx, cancel)

	filter := bson.M{"_id": bson.M{"$eq": mc.Id}}

	update := bson.M{
		"$set": bson.M{
			"name":      mc.Name,
			"symbol":    mc.Symbol,
			"votes":     mc.Votes,
			"createdat": mc.CreatedAT,
			"updatedat": mc.UpdatedAT,
		},
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	collection.FindOneAndDelete(context.Background(), filter)

	return err
}
