package model

import (
	"DesafioTecnico/misc"
	"fmt"

	"github.com/google/uuid"
)

type MoedaCripto struct {
	Id      string `json:"id" bson:"id"`
	Nome    string `json:"nome" bson:"nome"`
	Simbolo string `json:"simbolo" bson:"simbolo"`
	Votos   int    `json:"votos" bson:"votos"`
}

func (mc *MoedaCripto) CriarNovaCriptoMoeda() *MoedaCripto {
	misc.Limpatela()
	fmt.Print("CRIANDO NOVA MOEDA\n\n")
	mc.Id = uuid.NewString()
	fmt.Print("INFORME O NOME DA MOEDA: ")
	fmt.Scan(&mc.Nome)
	fmt.Print("INFORME O SIMBOLO DA MOEDA: ")
	fmt.Scan(&mc.Simbolo)
	mc.Votos = 0
	misc.Limpatela()
	return mc
}

func (mc *MoedaCripto) EditarCriptoMoeda(id string, lista []MoedaCripto) (indice int ,obj *MoedaCripto ) {
	
	pause := ""
	for i, item := range lista {
		if id == item.Id {
			misc.Limpatela()
			fmt.Print("TOKEN: ", item.Nome, " ENCONTRADO NA POSICAO: ",i,"\n\n")

			fmt.Print("NOVO NOME: ")
			fmt.Scan(&item.Nome)

			fmt.Print("NOVO SIMBOLO: ")
			fmt.Scan(&item.Simbolo)
			
			misc.Limpatela()
			return i, &item
		}
	}
	fmt.Printf("Moeda nao encontrada")
	fmt.Scan(&pause)
	misc.Limpatela()

	return 
}

func (mc *MoedaCripto) DeletarCriptoMoeda(id string, lista []MoedaCripto) (indice int, found bool) {
	pause := ""
	found = false
	
	for i, elemento := range(lista){
		if(id == elemento.Id){
			found = true
			return i, found;
		}
	} 
	
	misc.Limpatela()
	fmt.Print("MOEDA NAO ENCONTRADA")
	fmt.Print(indice)
	fmt.Scan(&pause)

	
	
	return
}

func (mc *MoedaCripto) UpVote(id string) {
	fmt.Print("Teste voteup nova moeda")
}

func (mc *MoedaCripto) DownVote(id string) {
	fmt.Print("Teste votedown nova moeda")
}

func (mc *MoedaCripto) ListarCriptoMoedas() {
	fmt.Print("Teste listartudo  moeda")
}
