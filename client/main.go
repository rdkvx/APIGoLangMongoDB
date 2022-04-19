package main

import (
	"DesafioTecnico/client/misc"
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
	//Moedas := []model.MoedaCripto{}
	for op {

		switch misc.MenuInicial() {
		case 1:
			services.CriarNovaCriptoMoedaAPI()
		case 2:
			misc.Limpatela()
			tempId := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)

			res, err := services.ListarUmaCriptoAPI(tempId)

			if err == nil {
				services.EditarCriptoMoedaAPI(res.Id, res)
			}
		case 3:
			misc.Limpatela()
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.ListarUmaCriptoAPI(tempID)
			if err == nil {
				services.DeletarCriptoMoedaAPI(res.Id)
			}
		case 4:
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.ListarUmaCriptoAPI(tempID)

			if err == nil {
				services.UpVoteAPI(res.Id, res)
			}
		case 5:
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.ListarUmaCriptoAPI(tempID)

			if err == nil {
				services.DownVoteAPI(res.Id, res)
			}
		case 6:
			err := services.ListarCriptoMoedasAPI()

			if err != nil {
				fmt.Println("ERRO AO LISTAR CRIPTOS")
				fmt.Scan(&pause)
			}
		case 0:
			op = false
		default:
			fmt.Print("Opcao Invalida")

		}
	}
}
