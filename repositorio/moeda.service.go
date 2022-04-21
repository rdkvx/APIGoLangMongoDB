package repositorio

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	pb "DesafioTecnico/proto"
	"DesafioTecnico/proto/conexaoservidor"
	m "DesafioTecnico/server/model"
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func CreateNewCryptoClient() {
	conexao, client, contexto, cancel := conexaoservidor.ConectaServidor()
	defer conexao.Close()
	defer cancel()

	pause := ""

	misc.CleanScreen()
	mc := m.MoedaCripto{}

	fmt.Print("CRIANDO NOVA MOEDA\n\n")

	fmt.Print("INFORME O NAME DA MOEDA: ")
	fmt.Scan(&mc.Name)

	fmt.Print("INFORME O SYMBOL DA MOEDA: ")
	fmt.Scan(&mc.Symbol)

	mc.CreatedAT = time.Now().Format("02/01/2006 15:04:45")

	misc.CleanScreen()

	response, err := client.CreateACrypto(contexto, &pb.RequestNewCrypto{Id: mc.Id, Name: mc.Name, Symbol: mc.Symbol, Createdat: mc.CreatedAT})
	if err != nil {
		log.Fatalf("ERRO AO GRAVAR CRIPTO: %v", err)
	} else {
		misc.CleanScreen()
		fmt.Println("CRIPTO GRAVADA COM SUCESSO!")
		fmt.Println("NAME: ", response.ResponseCripto.Name)
		fmt.Println("SYMBOL: ", response.ResponseCripto.GetSymbol())
		fmt.Println("CREATED AT: ", response.ResponseCripto.GetCreatedat())
		fmt.Scan(&pause)
		misc.CleanScreen()
	}
}

/* func EditarCriptoMoedaAPI(mc m.MoedaCripto) (m.MoedaCripto, error) {

	err := Update(mc)

	if err != nil {
		return mc, err
	}


	return mc, nil
}
*/
func EditingACryptoClient() {
	misc.CleanScreen()
	tempID := ""
	pause := ""

	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := SearchingACryptoAPI(tempID)
	if err != nil {
		misc.CleanScreen()
		fmt.Println("ID INVALIDO")
		fmt.Scan(&pause)
	} else {
		fmt.Print("\nNOVO NAME: ")
		fmt.Scan(&res.Name)

		fmt.Print("NOVO SYMBOL: ")
		fmt.Scan(&res.Symbol)

		res.UpdatedAT = time.Now().Format("02/01/2006 15:04:45")
		/*err := Update(res)
		if err != nil {
			log.Panic("ERRO AO EDITAR CRIPTO! ", err)
		}*/

		//response

		m.FeedbackEditarCriptoMoeda(res)
	}
}

func DeletingACryptoAPI(id string) error {
	err := Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func DeletingACryptoClient() {
	misc.CleanScreen()
	tempID := ""
	pause := ""

	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := SearchingACryptoAPI(tempID)
	if err != nil {
		misc.CleanScreen()
		fmt.Print("ID INVALIDO")
		fmt.Scan(&pause)
		misc.CleanScreen()
	} else {
		err = DeletingACryptoAPI(res.Id)
		if err != nil {
			misc.CleanScreen()
			fmt.Println("ERRO AO DELETAR CRIPTO! ", err)
			fmt.Scan(&pause)
			misc.CleanScreen()
		} else {
			misc.CleanScreen()
			fmt.Println("CRIPTO ", res.Name, " DELETADA COM SUCESSO")
			fmt.Scan(&pause)
			misc.CleanScreen()
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

	_, err = UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

	if err == nil {
		mc, err := Read(id)
		return mc, err
	}

	return m.MoedaCripto{}, err
}

func UpVoteClient() {
	tempID := ""
	pause := ""

	misc.CleanScreen()
	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := SearchingACryptoAPI(tempID)

	if err != nil {
		misc.CleanScreen()
		fmt.Println("ID INVALIDO!")
		fmt.Scan(&pause)
		misc.CleanScreen()

	} else {
		mc, err := UpVoteAPI(res.Id)

		if err == nil {
			misc.CleanScreen()
			fmt.Println("VOTE REGISTRED")
			fmt.Println("UPDATED INFOS!")
			fmt.Println("")
			fmt.Println("NAME : ", mc.Name)
			fmt.Println("VOTES: ", mc.Votes)
			fmt.Scan(&pause)
			misc.CleanScreen()
		} else {
			fmt.Println("ERROR REGISTERING VOTE!")
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

	_, err = UpdateOne(client, context.Background(), "desafiotecnico", "moedacripto", filter, update)

	if err == nil {
		mc, err := Read(id)
		return mc, err
	}

	return m.MoedaCripto{}, err
}

func DownVoteClient() {
	tempID := ""
	pause := ""

	misc.CleanScreen()
	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempID)

	res, err := SearchingACryptoAPI(tempID)

	if err != nil {
		misc.CleanScreen()
		fmt.Println("ID INVALIDO!")
		fmt.Scan(&pause)
		misc.CleanScreen()

	} else {
		mc, err := DownVoteAPI(res.Id)

		if err == nil {
			misc.CleanScreen()
			fmt.Println("VOTES REGISTERED")
			fmt.Println("UPDATED INFOS!")
			fmt.Println("")
			fmt.Println("NAME : ", mc.Name)
			fmt.Println("VOTES: ", mc.Votes)
			fmt.Scan(&pause)
			misc.CleanScreen()
		} else {
			fmt.Println("ERROR REGISTERING VOTE!")
			panic(err)
		}
	}
}

func SearchingACryptoAPI(id string) (m.MoedaCripto, error) {
	mc, err := Read(id)

	return mc, err
}

func ListCryptosAPI() (obj []m.MoedaCripto, err error) {
	//pause := ""

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

func ListCryptosClient() {
	pause := ""
	mc, err := ListCryptosAPI()

	if err != nil {
		fmt.Println("ERROR TRYING TO LIST CRYPTOS")
		fmt.Scan(&pause)
	} else {
		misc.CleanScreen()
		for i, elemento := range mc {
			fmt.Println("")
			fmt.Println("CRYPTO ", i+1)
			fmt.Println("NAME: ", elemento.Name)
			fmt.Println("SYMBOL: ", elemento.Symbol)
			fmt.Println("VOTOS: ", elemento.Votes)
		}
		fmt.Scan(&pause)
		misc.CleanScreen()
	}
}
