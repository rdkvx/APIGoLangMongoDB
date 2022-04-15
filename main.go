package main

import (
	"DesafioTecnico/misc"
	"DesafioTecnico/model"
	"fmt"
)

func main() {
	op := true
	pause := ""
	moedas := []model.MoedaCripto{}
	for op {

		switch misc.MenuInicial() {
		case 1:
			mc := model.MoedaCripto{}
			mc.CriarNovaCriptoMoeda()
			moedas = append(moedas, mc)
		case 2:
			misc.Limpatela()
			mc := model.MoedaCripto{}
			tempId := ""
			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)
			indice, obj, found := mc.EditarCriptoMoeda(tempId, moedas)
			if found {
				moedas[indice] = *obj
			}
		case 3:
			misc.Limpatela()

			mc := model.MoedaCripto{}
			tempId := ""

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)

			indice, found := mc.DeletarCriptoMoeda(tempId, moedas)
			if found {
				moedas[indice].Nome = ""
				moedas[indice].Id = ""
				moedas[indice].Simbolo = ""
				moedas[indice].Votos = 0
				fmt.Print("MOEDA DELETADA COM SUCESSO")
				fmt.Scan(&pause)
				misc.Limpatela()
			}
		case 4:
			tempID := ""
			mc := model.MoedaCripto{}

			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempID)

			voteup, found, indice := mc.UpVote(tempID, moedas)
			if found {
				misc.Limpatela()
				moedas[indice].Votos = voteup
				fmt.Print("VOTO REGISTRADO!\n")
				fmt.Println("MOEDA: ", moedas[indice].Nome)
				fmt.Println("VOTOS: ", moedas[indice].Votos)
				fmt.Scan(&pause)
			}

		case 5:
			tempId := ""
			fmt.Print("INFORME O ID DA MOEDA: ")
			fmt.Scan(&tempId)

			mc := model.MoedaCripto{}
			votedown, found, indice := mc.DownVote(tempId, moedas)
			if found {
				misc.Limpatela()
				moedas[indice].Votos = votedown
				fmt.Println("VOTO REGISTRADO!")
				fmt.Println("MOEDA: ", moedas[indice].Nome)
				fmt.Println("VOTOS: ", moedas[indice].Votos)
				fmt.Scan(&pause)
			}
		case 6:
			misc.Limpatela()
			mc := model.MoedaCripto{}
			mc.ListarCriptoMoedas(moedas)
			fmt.Scan(&pause)
			misc.Limpatela()
		case 0:
			op = false
		default:
			fmt.Print("Opcao Invalida")

		}
	}
}
