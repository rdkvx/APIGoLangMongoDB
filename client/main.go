package main

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/server/model"
	services "DesafioTecnico/server/services/moeda.service"
	"fmt"
)

func main() {
	// collection := client.Database(dataBase).Collection(col)
	// result, err = collection.Find(ctx, bson.D{})

	/*teste, _ := services.ListarUmaCriptoAPI("0375b319-85db-4f76-8ae7-5990a763fba1")
	fmt.Println(teste)*/

	op := true
	pause := ""
	Moedas := []model.MoedaCripto{}
	for op {

		switch misc.MenuInicial() {
		case 1:
			mc := model.MoedaCripto{}
			mc.CriarNovaCriptoMoeda()
			services.CriarNovaCriptoMoedaAPI(mc)
		case 2:
			misc.Limpatela()
			mc := model.MoedaCripto{}
			tempId := ""
			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)
			res, err := services.ListarUmaCriptoAPI(tempId)
			if (err != nil){
				fmt.Print("erro : ",err)
			}
			fmt.Print("resposta: ", res)
			fmt.Scan(&pause)
			obj, found := mc.EditarCriptoMoeda(tempId, Moedas)
			if found {
				//Moedas[indice] = *obj
				services.EditarCriptoMoedaAPI(tempId, *obj)
			}
		case 3:
			misc.Limpatela()

			mc := model.MoedaCripto{}
			tempId := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)

			indice, found := mc.DeletarCriptoMoeda(tempId, Moedas)
			if found {
				Moedas[indice].Nome = ""
				Moedas[indice].Id = ""
				Moedas[indice].Simbolo = ""
				Moedas[indice].Voto = 0
				fmt.Print("MOEDA DELETADA COM SUCESSO")
				fmt.Scan(&pause)
				misc.Limpatela()
			}
		case 4:
			tempID := ""
			mc := model.MoedaCripto{}

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			voteup, found, indice := mc.UpVote(tempID, Moedas)
			if found {
				misc.Limpatela()
				Moedas[indice].Voto = voteup
				fmt.Print("VOTO REGISTRADO!\n")
				fmt.Println("MOEDA: ", Moedas[indice].Nome)
				fmt.Println("Voto: ", Moedas[indice].Voto)
				fmt.Scan(&pause)
			}

		case 5:
			tempId := ""
			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)

			mc := model.MoedaCripto{}
			votedown, found, indice := mc.DownVote(tempId, Moedas)
			if found {
				misc.Limpatela()
				Moedas[indice].Voto = votedown
				fmt.Println("VOTO REGISTRADO!")
				fmt.Println("MOEDA: ", Moedas[indice].Nome)
				fmt.Println("Voto: ", Moedas[indice].Voto)
				fmt.Scan(&pause)
			}
		case 6:
			misc.Limpatela()
			mc := model.MoedaCripto{}
			mc.ListarCriptoMoedas(Moedas)
			fmt.Scan(&pause)
			misc.Limpatela()
		case 0:
			op = false
		default:
			fmt.Print("Opcao Invalida")

		}
	}
}
