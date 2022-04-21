package misc

import (
	"fmt"
	//"DesafioTecnico/server/model"
)

func Clear() {
	fmt.Println("\033[2J")
}

func SkipLine() {
	fmt.Print("\n\n")
}
