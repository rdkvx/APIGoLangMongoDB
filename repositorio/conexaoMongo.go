package repositorio

import (
	
	"fmt"
	

	
	"gopkg.in/mgo.v2"
)

var SessaoMongo *mgo.Session

func AbreSessao() (err error){
	SessaoMongo, err = mgo.Dial("mongodb+srv://rdkvx:Dragonforce123@cluster0.ntd1x.mongodb.net/test")
	if err != nil{
		fmt.Print("Erro ao abrir sessao")
	}
	return
}