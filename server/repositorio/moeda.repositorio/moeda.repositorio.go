package moeda_repositorio

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	"context"
	"log"

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
	pause := ""

	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur := collection.FindOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	//defer cur.Close(context.Background())

	var results = m.MoedaCripto{}

	/* 	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	} */

	if err = cur.Decode(&results); err != nil {
		misc.Limpatela()
		fmt.Println("ID INVALIDO")
		fmt.Scan(&pause)
		/* log.Panic(err) */
		return
	}
	//mc = results[0]
	fmt.Print("MOEDA: ", results.Nome)

	return results, err
}

func Update(id string, mc m.MoedaCripto) error {
	pause := ""

	// get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer database.Close(client, ctx, cancel)

	// filter object is used to select a single
	// document matching that matches.
	//filter := bson.M{"moedacripto": bson.M{"_id": id},}
	filter := bson.M{"_id": bson.M{"$eq": id}}

	// The field of the document that need to updated.
	update := bson.M{
		"$set": bson.M{"name": mc.Nome,
			"symbol":    mc.Simbolo,
			"updatedat": mc.UpdatedAT,
		},
	}

	// Returns result of updated document and a error.
	result, err := database.UpdateOne(client, context.Background(), "desafiotecnico",
		"moedacripto", filter, update)

	// handle error
	if err != nil {
		panic(err)
	}

	misc.Limpatela()
	fmt.Println("MOEDA ATUALIZADA COM SUCESSO")
	fmt.Println("INFORMACOES ATUALIZADAS!")
	fmt.Println("")
	fmt.Println("NOME: ", mc.Nome)
	fmt.Println("SIMBOLO: ", mc.Simbolo)
	fmt.Println("DATA DE CRIACAO: ", mc.CreatedAT)
	fmt.Println("DATA DA ATUALIZACAO: ", mc.UpdatedAT)
	fmt.Println("TOTAL DE DOCUMENTOS ATUALIZADOS: ", result.ModifiedCount)
	fmt.Scan(&pause)
	misc.Limpatela()

	return nil
}

func Delete(id string) error {
	pause := ""

	client, _, _, err := database.Connect(database.Uri())

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	filter := bson.M{"_id": id}

	cur := collection.FindOneAndDelete(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var results = m.MoedaCripto{}

	if err = cur.Decode(&results); err != nil {
		misc.Limpatela()
		fmt.Println("ID INVALIDO")
		fmt.Scan(&pause)
		return err
	}

	fmt.Println(" REMOVIDA COM SUCESSO")
	fmt.Scan(&pause)

	return err
}
