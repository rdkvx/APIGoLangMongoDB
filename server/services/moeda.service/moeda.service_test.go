package services_test

import (
	"DesafioTecnico/server/model"
	services "DesafioTecnico/server/services/moeda.service"
	"testing"
)

func TestCriarNovaCriptoMoedaAPI(t *testing.T) {
	moeda := model.MoedaCripto{
		Id: "1",
		Nome: "rodriCoin",
		Simbolo: "RC",
		Voto: 1,
	}
	err := services.CriarNovaCriptoMoedaAPI(moeda)

	if err != nil{
		t.Error("Erro ao inserir usuario")
		t.Fail()
	}else{
		t.Log("Tudo ok!")
	}
}


func TestEditarCriptoMoedaAPI(t *testing.T){

}

func TestDeletarCriptoMoedaAPI(t *testing.T){

}



func TestListarUmaCriptoAPI(t *testing.T){

}

func TestListarCriptoMoedasAPI(t *testing.T){

}