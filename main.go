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
			indice, obj := mc.EditarCriptoMoeda(tempId, moedas)
			moedas[indice] = *obj
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
			mc := model.MoedaCripto{}
			mc.UpVote(mc.Id)
		case 5:
			mc := model.MoedaCripto{}
			mc.DownVote(mc.Id)
		case 6:
			misc.Limpatela()
			//mc :=model.MoedaCripto{}
			fmt.Print(moedas)
			fmt.Scan(&pause)
			misc.Limpatela()
		case 0:
			op = false
		default:
			fmt.Print("Opcao Invalida")

		}
	}
}
