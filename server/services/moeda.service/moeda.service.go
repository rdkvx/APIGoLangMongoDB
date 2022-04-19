package services

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	m "DesafioTecnico/server/model"
	moeda_repositorio "DesafioTecnico/server/repositorio/moeda.repositorio"
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/google/uuid"
)

func CriarNovaCriptoMoedaAPI() error {
	mc := m.MoedaCripto{}

	misc.Limpatela()
	fmt.Print("CRIANDO NOVA MOEDA\n\n")
	mc.Id = uuid.NewString()
	fmt.Print("INFORME O NOME DA MOEDA: ")
	fmt.Scan(&mc.Nome)
	fmt.Print("INFORME O SIMBOLO DA MOEDA: ")
	fmt.Scan(&mc.Simbolo)
	mc.Voto = 0
	mc.CreatedAT = time.Now().Format("02/01/2006 15:04:45")
	misc.Limpatela()

	err := moeda_repositorio.Create(mc)

	if err != nil {
		return err
	}

	return nil
}

func EditarCriptoMoedaAPI(id string, mc m.MoedaCripto) error {
	fmt.Print("\nNOVO NOME: ")
	fmt.Scan(&mc.Nome)

	fmt.Print("NOVO SIMBOLO: ")
	fmt.Scan(&mc.Simbolo)

	mc.UpdatedAT = time.Now().Format("02/01/2006 15:04:45")

	err := moeda_repositorio.Update(id, mc)

	if err != nil {
		return err
	}
	return nil
}

func DeletarCriptoMoedaAPI(id string) error {
	err := moeda_repositorio.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func UpVoteAPI(id string, mc m.MoedaCripto) error {
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
	filter := bson.M{"_id": id}

	// The field of the document that need to updated.
	update := bson.M{
		"$inc": bson.M{"votes": 1},
	}

	// Returns result of updated document and a error.
	_, err = database.UpdateOne(client, context.Background(), "desafiotecnico",
		"moedacripto", filter, update)

	// handle error
	if err != nil {
		panic(err)
	}

	misc.Limpatela()
	fmt.Println("VOTO COMPUTADO")
	fmt.Println("INFORMACOES ATUALIZADAS!")
	fmt.Println("")
	fmt.Println("NOME : ", mc.Nome)
	fmt.Println("VOTOS: ", mc.Voto+1)
	//fmt.Println("TOTAL DE DOCUMENTOS ATUALIZADOS: ", result.ModifiedCount)
	fmt.Scan(&pause)
	misc.Limpatela()

	return err

}

func DownVoteAPI(id string, mc m.MoedaCripto) error {
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
		"$inc": bson.M{"votes": -1},
	}

	if mc.Voto > 0 {
		// Returns result of updated document and a error.
		result, err := database.UpdateOne(client, context.Background(), "desafiotecnico",
			"moedacripto", filter, update)

		// handle error
		if err != nil {
			panic(err)
		}

		misc.Limpatela()
		fmt.Println("VOTO COMPUTADO")
		fmt.Println("INFORMACOES ATUALIZADAS!")
		fmt.Println("")
		fmt.Println("NOME : ", mc.Nome)
		fmt.Println("VOTOS: ", mc.Voto-1)
		fmt.Println("TOTAL DE DOCUMENTOS ATUALIZADOS: ", result.ModifiedCount)
		fmt.Scan(&pause)
		misc.Limpatela()
		return err
	}

	misc.Limpatela()
	fmt.Println("A MOEDA NAO POSSUI VOTO")
	fmt.Scan(&pause)
	misc.Limpatela()

	return err
}

func ListarUmaCriptoAPI(id string) (m.MoedaCripto, error) {
	mc, err := moeda_repositorio.Read(id)

	if err != nil {
		return mc, err
	}
	return mc, err
}

func ListarCriptoMoedasAPI() (err error) {
	pause := ""

	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("Erro ao conectar na base: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	//filter := bson.M{"_id": id}

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//defer cur.Close(context.Background())

	var results = []m.MoedaCripto{}

	/* 	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	} */

	if err = cur.All(context.Background(), &results); err != nil {
		misc.Limpatela()
		fmt.Println("ERRO MOEDA SERVICE")
		fmt.Scan(&pause)
		/* log.Panic(err) */
		return err
	}

	misc.Limpatela()
	for i, elemento := range results {
		fmt.Println("")
		fmt.Println("MOEDA ", i+1)
		fmt.Println("NOME: ", elemento.Nome)
		fmt.Println("SIMBOLO: ", elemento.Simbolo)
		fmt.Println("VOTOS: ", elemento.Voto)
	}

	fmt.Scan(&pause)

	return err
}
