package misc

import (
	"fmt"
	//"DesafioTecnico/server/model"
)

func MainMenu() int {
	op := 0

	fmt.Println("\nDESAFIO TECNICO")
	fmt.Println("[1] - CRIA NOVA MOEDA")
	fmt.Println("[2] - EDITAR MOEDA")
	fmt.Println("[3] - DELETAR MOEDA")
	fmt.Println("[4] - VOTEUP")
	fmt.Println("[5] - VOTEDOWN")
	fmt.Println("[6] - LISTAR MOEDAS")
	fmt.Println("[0] - SAIR")
	fmt.Print("\nOPCAO: ")
	fmt.Scan(&op)

	return op

}

func Clear() {
	fmt.Println("\033[2J")
}

func SkipLine() {
	fmt.Print("\n\n")
}
