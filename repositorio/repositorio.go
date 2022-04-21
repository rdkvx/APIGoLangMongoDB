package repositorio

import (
	"DesafioTecnico/database"
	"context"
	"log"

	m "DesafioTecnico/server/model"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Create(mc m.MoedaCripto) error {

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

func Read(id string) (mc m.MoedaCripto, err error) {

	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur := collection.FindOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var results = m.MoedaCripto{}

	if err = cur.Decode(&results); err != nil {
		return results, err
	}

	return results, err
}

func ReadAll() (obj []m.MoedaCripto, err error) {
	
	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var results = []m.MoedaCripto{}

	if err = cur.All(context.Background(), &results); err != nil {
		return []m.MoedaCripto{}, err
	}

	return results, err
}

func Update(mc m.MoedaCripto) error {

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	filter := bson.M{"_id": bson.M{"$eq": mc.Id}}

	update := bson.M{
		"$set": bson.M{"name": mc.Name,
			"symbol":    mc.Symbol,
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

	var results = m.MoedaCripto{}

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
