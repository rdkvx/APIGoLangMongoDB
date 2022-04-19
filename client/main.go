package main

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/server/model"
	services "DesafioTecnico/server/services/moeda.service"
	"fmt"
	//"log"
)

func main() {

	op := true
	pause := ""
	for op {

		switch misc.MenuInicial() {
		case 1:
			mc := model.MoedaCripto{}
			mc.CriarNovaCriptoMoeda()
			services.CriarNovaCriptoMoedaAPI(mc)
		case 2:
			services.EditaCriptoMoedaClient()
		case 3:
			misc.Limpatela()
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.BuscarUmaCriptoAPI(tempID)
			if err == nil {
				services.DeletarCriptoMoedaAPI(res.Id)
			}
		case 4:
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.BuscarUmaCriptoAPI(tempID)

			if err == nil {
				services.UpVoteAPI(res.Id)
			}
		case 5:
			tempID := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			res, err := services.BuscarUmaCriptoAPI(tempID)

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
