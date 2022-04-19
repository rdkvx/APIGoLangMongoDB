package services

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	m "DesafioTecnico/server/model"
	mr "DesafioTecnico/server/repositorio/moeda.repositorio"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"gopkg.in/mgo.v2/bson"
)

func CriarNovaCriptoMoedaAPI(mc m.MoedaCripto) error {

	err := mr.Create(mc)

	if err != nil {
		return err
	}
	return nil
}

func CriarNovaCriptoMoedaClient() {
	pause := ""

	misc.Limpatela()
	mc := m.MoedaCripto{}

	fmt.Print("CRIANDO NOVA MOEDA\n\n")
	mc.Id = uuid.NewString()

	fmt.Print("INFORME O NOME DA MOEDA: ")
	fmt.Scan(&mc.Nome)

	fmt.Print("INFORME O SIMBOLO DA MOEDA: ")
	fmt.Scan(&mc.Simbolo)

	mc.Voto = 0
	mc.CreatedAT = time.Now().Format("02/01/2006 15:04:45")

	misc.Limpatela()

	err := CriarNovaCriptoMoedaAPI(mc)
	if err == nil {
		misc.Limpatela()
		fmt.Println("MOEDA CRIADA COM SUCESSO")
		fmt.Scan(&pause)
		misc.Limpatela()
	} else {
		misc.Limpatela()
		fmt.Println("ERRO AO CRIAR MOEDA!")
		fmt.Scan(&pause)
		misc.Limpatela()
	}
}

func EditarCriptoMoedaAPI(mc m.MoedaCripto) (m.MoedaCripto, error) {

	err := mr.Update(mc)

	if err != nil {
		return mc, err
	}

	mc.UpdatedAT = time.Now().Format("02/01/2006 15:04:45")
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
		obj, err := EditarCriptoMoedaAPI(res)
		if err != nil {
			log.Panic("ERRO AO EDITAR CRIPTO! ", err)
		}
		m.FeedbackEditarCriptoMoeda(obj)
	}
}

func DeletarCriptoMoedaAPI(id string) error {
	err := mr.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func DeletarCriptoMoedaClient() {
	misc.Limpatela()
	tempID := ""
	pause := ""

	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := BuscarUmaCriptoAPI(tempID)
	if err != nil {
		misc.Limpatela()
		fmt.Print("ID INVALIDO")
		fmt.Scan(&pause)
		misc.Limpatela()
	} else {
		err = DeletarCriptoMoedaAPI(res.Id)
		if err != nil {
			misc.Limpatela()
			fmt.Println("ERRO AO DELETAR CRIPTO! ", err)
			fmt.Scan(&pause)
			misc.Limpatela()
		} else {
			misc.Limpatela()
			fmt.Println("CRIPTO ", res.Nome, " DELETADA COM SUCESSO")
			fmt.Scan(&pause)
			misc.Limpatela()
		}
	}

}

func UpVoteAPI(id string) (m.MoedaCripto, error) {

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

	if err == nil {
		mc, err := mr.Read(id)
		return mc, err
	}

	return m.MoedaCripto{}, err
}

func UpVoteClient() {
	tempID := ""
	pause := ""

	misc.Limpatela()
	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := BuscarUmaCriptoAPI(tempID)

	if err != nil {
		misc.Limpatela()
		fmt.Println("ID INVALIDO!")
		fmt.Scan(&pause)
		misc.Limpatela()

	} else {
		mc, err := UpVoteAPI(res.Id)

		if err == nil {
			misc.Limpatela()
			fmt.Println("VOTO COMPUTADO")
			fmt.Println("INFORMACOES ATUALIZADAS!")
			fmt.Println("")
			fmt.Println("NOME : ", mc.Nome)
			fmt.Println("VOTOS: ", mc.Voto)
			fmt.Scan(&pause)
			misc.Limpatela()
		} else {
			fmt.Println("ERRO AO REGISTRAR VOTO!!!")
			panic(err)
		}
	}
}

func DownVoteAPI(id string) (m.MoedaCripto, error) {
	client, ctx, cancel, err := database.Connect(database.Uri())
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	filter := bson.M{"_id": id}

	update := bson.M{
		"$inc": bson.M{"votes": -1},
	}

	_, err = mr.UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

	if err == nil {
		mc, err := mr.Read(id)
		return mc, err
	}

	return m.MoedaCripto{}, err
}

func DownVoteClient() {
	tempID := ""
	pause := ""

	misc.Limpatela()
	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := BuscarUmaCriptoAPI(tempID)

	if err != nil {
		misc.Limpatela()
		fmt.Println("ID INVALIDO!")
		fmt.Scan(&pause)
		misc.Limpatela()

	} else {
		mc, err := DownVoteAPI(res.Id)

		if err == nil {
			misc.Limpatela()
			fmt.Println("VOTO COMPUTADO")
			fmt.Println("INFORMACOES ATUALIZADAS!")
			fmt.Println("")
			fmt.Println("NOME : ", mc.Nome)
			fmt.Println("VOTOS: ", mc.Voto)
			fmt.Scan(&pause)
			misc.Limpatela()
		} else {
			fmt.Println("ERRO AO REGISTRAR VOTO!!!")
			panic(err)
		}
	}
}

func BuscarUmaCriptoAPI(id string) (m.MoedaCripto, error) {
	mc, err := mr.Read(id)

	return mc, err
}

func ListarCriptoMoedasAPI() (obj []m.MoedaCripto,err error) {
	//pause := ""

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
		return []m.MoedaCripto{}, err
	}

	return results, err
}

func ListarCriptoMoedasClient() {
	pause := ""
	mc,err := ListarCriptoMoedasAPI()

	if err != nil {
		fmt.Println("ERRO AO LISTAR CRIPTOS")
		fmt.Scan(&pause)
	}else{
		misc.Limpatela()
		for i, elemento := range mc {
			fmt.Println("")
			fmt.Println("MOEDA ", i+1)
			fmt.Println("NOME: ", elemento.Nome)
			fmt.Println("SIMBOLO: ", elemento.Simbolo)
			fmt.Println("VOTOS: ", elemento.Voto)
		}
		fmt.Scan(&pause)
		misc.Limpatela()
	}
}
