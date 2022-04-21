package model

import (
	"DesafioTecnico/client/misc"
	"fmt"
)

type MoedaCripto struct {
	Id        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Symbol    string `json:"symbol" bson:"symbol"`
	Votes     int    `json:"votes" bson:"votes"`
	CreatedAT string `json:"createdat" bson:"createdat"`
	UpdatedAT string `json:"updatedat" bson:"updatedat"`
}

func FeedbackEditarCriptoMoeda(mc MoedaCripto) {
	pause := ""
	misc.CleanScreen()
	fmt.Println("CRYPTO UPDATED SUCCESSFULLY")
	fmt.Println("UPDATED INFOS!")
	fmt.Println("")
	fmt.Println("NAME: ", mc.Name)
	fmt.Println("SYMBOL: ", mc.Symbol)
	fmt.Println("CREATION DATE: ", mc.CreatedAT)
	fmt.Println("UPDATED INFO DATE: ", mc.UpdatedAT)
	fmt.Scan(&pause)
	misc.CleanScreen()
}
