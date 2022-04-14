package model

import (
	"github.com/google/uuid"
)

type MoedaCripto struct {
	Id      uuid.UUID `json:"id" bson:"id"`
	Nome    string `json:"nome" bson:"nome"`
	Simbolo string `json:"simbolo" bson:"simbolo"`
	Votos   int    `json:"votos" bson:"votos"`
}

func (mc *MoedaCripto) CriarNovaCriptoMoeda(nome string, simbolo string, votos int){
	mc.Id = uuid.New();
	
	mc.Nome = nome;
	mc.Simbolo = simbolo;
	mc.Votos = votos
}

func (mc *MoedaCripto) EditarCriptoMoeda(id uuid.UUID){

}

func (mc *MoedaCripto) DeletarCriptoMoeda(id uuid.UUID){

}

func (mc *MoedaCripto) UpVote(id uuid.UUID){

}

func (mc *MoedaCripto) DownVote(id uuid.UUID){

}

func (mc *MoedaCripto) ListarCriptoMoedas(){

}