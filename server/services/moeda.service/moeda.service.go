package services

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	"DesafioTecnico/server/model"
	m "DesafioTecnico/server/model"
	mr "DesafioTecnico/server/repositorio/moeda.repositorio"
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
	/* "github.com/google/uuid" */)

func CriarNovaCriptoMoedaAPI(mc m.MoedaCripto) error {

	err := mr.Create(mc)

	if err != nil {
		return err
	}
	return nil
}

func EditarCriptoMoedaAPI(id string, mc m.MoedaCripto) (m.MoedaCripto, error) {
	
	mc.UpdatedAT = time.Now().Format("02/01/2006 15:04:45")

	err := mr.Update(id, mc)

	if err != nil {
		return mc, err
	}
	return mc, nil
}

func EditaCriptoMoedaClient() {
	misc.Limpatela()
	tempID := ""
	pause := ""

	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := BuscarUmaCriptoAPI(tempID)
	if err != nil {
		misc.Limpatela()
		fmt.Println("ID INVALIDO")
		fmt.Scan(&pause)
	} else {
		fmt.Print("\nNOVO NOME: ")
		fmt.Scan(&res.Nome)

		fmt.Print("NOVO SIMBOLO: ")
		fmt.Scan(&res.Simbolo)
		obj, err := EditarCriptoMoedaAPI(res.Id, res)
		if err != nil {
			log.Panic("ERRO AO EDITAR CRIPTO! ", err)
		}
		model.FeedbackEditarCriptoMoeda(obj)
	}
	
}

func DeletarCriptoMoedaAPI(id string) error {
	err := mr.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func UpVoteAPI(id string) error {
	pause := ""

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	filter := bson.M{"_id": id}

	update := bson.M{
		"$inc": bson.M{"votes": 1},
	}

	_, err = mr.UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

	if err != nil {
		fmt.Println("ERRO AO REGISTRAR VOTO!!!")
		panic(err)
	}

	mc, err := mr.Read(id)

	misc.Limpatela()
	fmt.Println("VOTO COMPUTADO")
	fmt.Println("INFORMACOES ATUALIZADAS!")
	fmt.Println("")
	fmt.Println("NOME : ", mc.Nome)
	fmt.Println("VOTOS: ", mc.Voto)
	fmt.Scan(&pause)
	misc.Limpatela()

	return err
}

func DownVoteAPI(id string, mc m.MoedaCripto) error {
	pause := ""

	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	filter := bson.M{"_id": bson.M{"$eq": id}}

	update := bson.M{
		"$inc": bson.M{"votes": -1},
	}

	if mc.Voto > 0 {
		result, err := mr.UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

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

func BuscarUmaCriptoAPI(id string) (m.MoedaCripto, error) {
	mc, err := mr.Read(id)

	return mc, err
}

func ListarCriptoMoedasAPI() (err error) {
	pause := ""

	client, _, _, err := database.Connect(database.Uri())

	if err != nil {
		fmt.Println("Erro ao conectar na base: ", err)
	}

	collection := client.Database("desafiotecnico").Collection("moedacripto")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var results = []m.MoedaCripto{}

	if err = cur.All(context.Background(), &results); err != nil {
		misc.Limpatela()
		fmt.Println("ERRO MOEDA SERVICE")
		fmt.Scan(&pause)
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
