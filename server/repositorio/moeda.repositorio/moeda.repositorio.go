package moeda_repositorio

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	"context"
	"log"

	//"DesafioTecnico/server/model"
	//"DesafioTecnico/server/model"
	"DesafioTecnico/server/model"
	m "DesafioTecnico/server/model"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

/* func Uri() (conexao string) {
	return "mongodb+srv://rdkvx:Dragonforce123@cluster0.ntd1x.mongodb.net/desafiotecnico?retryWrites=true&w=majority"
} */

func Create(mc m.MoedaCripto) error {
	pause := ""

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	insertOneResult, err := database.InsertOne(client, ctx, "desafiotecnico", "moedacripto", mc)

	if err != nil {
		fmt.Print("Erro ao inserir, erro: ", err.Error())
	}

	// Release resource when the main
	// function is returned.
	defer database.Close(client, ctx, cancel)

	misc.Limpatela()
	fmt.Println("Inserts realizados")
	fmt.Println(insertOneResult)
	fmt.Println("MOEDA CRIADA COM SUCESSO")
	fmt.Scan(&pause)
	misc.Limpatela()

	return nil
}

func Read(id string) (mc m.MoedaCripto, err error) {
	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	pause := ""
	var results = []model.MoedaCripto{}

	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)

	}
	fmt.Print("moeda repositorio results: ", results)
	fmt.Scan(&pause)
	mc = results[0]

	return mc, err
}

func Update(id string, mc m.MoedaCripto) error {
	// get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer database.Close(client, ctx, cancel)

	// filter object is used to select a single
	// document matching that matches.
	filter := bson.D{
		{"moedacripto", bson.D{{"id", id}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{"$set", bson.D{
			{"name", mc.Nome},
			{"symbol", mc.Simbolo},
			{"updatedat", mc.UpdatedAT},
		}},
	}

	// Returns result of updated document and a error.
	result, err := database.UpdateOne(client, ctx, "desafiotecnico",
		"moedacripto", filter, update)

	// handle error
	if err != nil {
		panic(err)
	}

	// print count of documents that affected
	fmt.Println("update single document")
	fmt.Println(result.ModifiedCount)

	/* filter = bson.D{
		{"computer", bson.D{{"$lt", 100}}},
	}
	update = bson.D{
		{"$set", bson.D{
			{"computer", 100},
		}},
	}

	// Returns result of updated document and a error.
	result, err = Update(client, ctx, "desafiotecnico",
		"moedacripto", filter, update) */

	return nil
}

func Delete(id string) error {
	return nil
}
