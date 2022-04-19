package model

import (
	"DesafioTecnico/client/misc"
	"fmt"
)

type MoedaCripto struct {
	Id        string `json:"id" bson:"_id"`
	Nome      string `json:"name" bson:"name"`
	Simbolo   string `json:"symbol" bson:"symbol"`
	Voto      int    `json:"votes" bson:"votes"`
	CreatedAT string `json:"createdat" bson:"createdat"`
	UpdatedAT string `json:"updatedat" bson:"updatedat"`
}

func FeedbackEditarCriptoMoeda(mc MoedaCripto) {
	pause := ""
	misc.Limpatela()
	fmt.Println("MOEDA ATUALIZADA COM SUCESSO")
	fmt.Println("INFORMACOES ATUALIZADAS!")
	fmt.Println("")
	fmt.Println("NOME: ", mc.Nome)
	fmt.Println("SIMBOLO: ", mc.Simbolo)
	fmt.Println("DATA DE CRIACAO: ", mc.CreatedAT)
	fmt.Println("DATA DA ATUALIZACAO: ", mc.UpdatedAT)
	fmt.Scan(&pause)
	misc.Limpatela()
}
