package main

import (
	"DesafioTecnico/client/misc"
	"fmt"
)

func main() {
	pause := ""
	op := true
	for op {

		switch misc.MenuInicial() {
		case 1:
		// 	repo.CreateNewCryptoClient()
		// case 2:
		// 	repo.EditingACryptoClient()
		// case 3:
		// 	repo.DeletingACryptoClient()
		// case 4:
		// 	repo.UpVoteClient()
		// case 5:
		// 	repo.DownVoteClient()
		// case 6:
		// 	repo.ListCryptosClient()
		case 0:
			op = false
		default:
			misc.CleanScreen()
			fmt.Print("Opcao Invalida")
			fmt.Scan(&pause)
			misc.CleanScreen()
		}
	}
}
