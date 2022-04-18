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

//lista de moedas
type Moedas []MoedaCripto

func (mc *MoedaCripto) UpVote(id string, lista []MoedaCripto) (voteup int, found bool, indice int) {
	pause := ""
	found = false

	for indice, elemento := range lista {
		if id == elemento.Id {
			found = true
			voteup := elemento.Voto + 1

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
			if elemento.Voto > 0 {
				found = true
				votedown = lista[i].Voto - 1
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
		fmt.Print("VOTOS: ", elemento.Voto)
		misc.PulaLinha()
	}
}
