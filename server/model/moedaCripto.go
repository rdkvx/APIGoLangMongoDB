package model

import (
	"DesafioTecnico/client/misc"
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

func (mc *MoedaCripto) EditarCriptoMoeda(id string, lista []MoedaCripto) (indice int, obj *MoedaCripto, found bool) {
	found = false
	pause := ""
	for i, item := range lista {
		if id == item.Id {
			misc.Limpatela()
			fmt.Print("TOKEN: ", item.Nome, " ENCONTRADO NA POSICAO: ", i, "\n\n")

			fmt.Print("NOVO NOME: ")
			fmt.Scan(&item.Nome)

			fmt.Print("NOVO SIMBOLO: ")
			fmt.Scan(&item.Simbolo)

			found = true
			misc.Limpatela()
			return i, &item, found
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

	for i, elemento := range lista {
		if id == elemento.Id {
			found = true
			return i, found
		}
	}

	misc.Limpatela()
	fmt.Print("MOEDA NAO ENCONTRADA")
	fmt.Print(indice)
	fmt.Scan(&pause)

	return
}

func (mc *MoedaCripto) UpVote(id string, lista []MoedaCripto) (voteup int, found bool, indice int) {
	pause := ""
	found = false

	for indice, elemento := range lista {
		if id == elemento.Id {
			found = true
			voteup := elemento.Votos + 1

			return voteup, found, indice
		}
	}

	misc.Limpatela()
	fmt.Print("MOEDA NAO ENCONTRADA")
	fmt.Scan(&pause)

	return
}

func (mc *MoedaCripto) DownVote(id string, lista []MoedaCripto) (votedown int, found bool, indice int) {
	found = false
	pause := ""
	for i, elemento := range lista {
		if id == elemento.Id {
			if elemento.Votos > 0 {
				found = true
				votedown = lista[i].Votos - 1
				return votedown, found, i
			} else {
				misc.Limpatela()
				fmt.Print("A MOEDA INFORMADA NAO POSSUI VOTOS REGISTRADOS")
				fmt.Scan(&pause)
				misc.Limpatela()
				return
			}
		}
	}

	misc.Limpatela()
	fmt.Print("MOEDA NAO ENCONTRADA")
	fmt.Scan(&pause)

	return
}

func (mc *MoedaCripto) ListarCriptoMoedas(lista []MoedaCripto) {
	fmt.Print("TODAS AS MOEDAS REGISTRADAS\n\n")
	for i, elemento := range lista {
		if i == 0 {
			fmt.Println("TOTAL DE MOEDAS ENCONTRADAS: ", len(lista))
			misc.PulaLinha()
		}
		fmt.Println("MOEDA NUMERO: ", i+1)
		fmt.Println("ID: ", elemento.Id)
		fmt.Println("MOEDA: ", elemento.Nome)
		fmt.Println("SIMBOLO: ", elemento.Simbolo)
		fmt.Print("VOTOS: ", elemento.Votos)
		misc.PulaLinha()
	}
}
