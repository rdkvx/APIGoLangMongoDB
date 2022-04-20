package main

import (
	"DesafioTecnico/client/misc"
	repo "DesafioTecnico/repositorio"
	"fmt"
)

func main() {
	pause := ""
	op := true
	for op {

		switch misc.MenuInicial() {
		case 1:
			repo.CriarNovaCriptoMoedaClient()
		case 2:
			repo.EditaCriptoMoedaClient()
		case 3:
			repo.DeletarCriptoMoedaClient()
		case 4:
			repo.UpVoteClient()
		case 5:
			repo.DownVoteClient()
		case 6:
			repo.ListarCriptoMoedasClient()
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
