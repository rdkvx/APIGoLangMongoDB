package main

import (
	"DesafioTecnico/client/misc"
	services "DesafioTecnico/server/services/moeda.service"
	"fmt"
	
)

func main() {
	pause := ""
	op := true
	for op {

		switch misc.MenuInicial() {
		case 1:
			services.CriarNovaCriptoMoedaClient()
		case 2:
			services.EditaCriptoMoedaClient()
		case 3:
			services.DeletarCriptoMoedaClient()
		case 4:
			services.UpVoteClient()
		case 5:
			services.DownVoteClient()
		case 6:
			services.BuscarUmaCriptoClient()
		case 0:
			op = false
		default:
			misc.Limpatela()
			fmt.Print("Opcao Invalida")
			fmt.Scan(&pause)
			misc.Limpatela()
		}
	}
}
