package model

import (
	"DesafioTecnico/client/misc"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MoedaCripto struct {
	Id        string `json:"id" bson:"_id"`
	Nome      string `json:"name" bson:"name"`
	Simbolo   string `json:"symbol" bson:"symbol"`
	Voto      int    `json:"votes" bson:"votes"`
	CreatedAT string `json:"createdat" bson:"createdat"`
	UpdatedAT string `json:"updatedat" bson:"updatedat"`
}

func (mc *MoedaCripto) CriarNovaCriptoMoeda() *MoedaCripto {
	misc.Limpatela()
	fmt.Print("CRIANDO NOVA MOEDA\n\n")
	mc.Id = uuid.NewString()
	fmt.Print("INFORME O NOME DA MOEDA: ")
	fmt.Scan(&mc.Nome)
	fmt.Print("INFORME O SIMBOLO DA MOEDA: ")
	fmt.Scan(&mc.Simbolo)
	mc.Voto = 0
	mc.CreatedAT = time.Now().Format("02/01/2006 15:04:45")
	misc.Limpatela()
	return mc
}

/* func EditaCriptoMoeda() (id string) {
	misc.Limpatela()
	tempId := ""

	fmt.Print("INFORME O ID DA MOEDA: ")
	fmt.Scan(&tempId)

	return tempId
} */

func FeedbackEditarCriptoMoeda(mc MoedaCripto){
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
