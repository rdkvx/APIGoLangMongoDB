package utils

import (
	"fmt"
)

func Clear() {
	fmt.Println("\033[2J")
}

func SkipLine() {
	fmt.Print("\n\n")
}
