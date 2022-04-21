package repositorio

import (
	"DesafioTecnico/database"
	"context"
	"log"

	m "DesafioTecnico/server/model"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func Create(mc m.CryptoCurrency) error {

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	_, err = InsertOne(client, ctx, "desafiotecnico", "moedacripto", mc)

	if err != nil {
		fmt.Print("Erro ao inserir, erro: ", err.Error())
	}

	defer database.Close(client, ctx, cancel)

	return err
}

func Read(id string) (mc m.CryptoCurrency, err error) {

	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur := collection.FindOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
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

	if sortParam == "" && !ascending {
		//retornar a lista no formato padrao
		cur, err = collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
	} else { //lista com ordenacao customizada
		if sortParam != "" { //se tem parametro valido para ordenar por essa coluna correspondente
			if ascending { //se ascending = true, trazer no padraao (ascendente)
				// Sort by `sortParam` field ascending
				findOptions.SetSort(bson.M{sortParam: 1})
				cur, err = collection.Find(context.Background(), bson.M{}, findOptions)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				// Sort by `sortParam` field descending
				findOptions.SetSort(bson.M{sortParam: -1})
				cur, err = collection.Find(context.Background(), bson.M{}, findOptions)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	var results = []m.CryptoCurrency{}

	if err = cur.All(context.Background(), &results); err != nil {
		return []m.CryptoCurrency{}, err
	}

	return results, err
}

func Update(mc m.CryptoCurrency) error {

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
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

	_, err = UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

	if err != nil {
		panic(err)
	}

	return nil
}

func Delete(id string) error {

	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur := collection.FindOneAndDelete(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var results = m.CryptoCurrency{}

	if err = cur.Decode(&results); err != nil {

		return err
	}

	return err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		fmt.Print("Erro ao inserir: ", err)
	}

	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateOne(ctx, filter, update)

	return
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, filter, option interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.Find(context.Background(), filter)
	return
}
