package moeda_repositorio

import (
	"DesafioTecnico/client/misc"
	"DesafioTecnico/database"
	//"DesafioTecnico/server/model"
	m "DesafioTecnico/server/model"
	"fmt"
)

func Create(mc m.MoedaCripto) error {
	client, ctx, cancel, err := database.Connect("mongodb+srv://rdkvx:Dragonforce123@cluster0.ntd1x.mongodb.net/desafiotecnico?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

	insertOneResult, err := database.InsertOne(client, ctx, "desafiotecnico", "moedacripto", mc)
	

	if err != nil {
		fmt.Print("Erro ao inserir, erro: ", err.Error())
	}

	// Release resource when the main
	// function is returned.
	defer database.Close(client, ctx, cancel)

	misc.Limpatela()
	fmt.Println("Inserts realizados")
	fmt.Println(insertOneResult.InsertedID)
	fmt.Println("MOEDA CRIADA COM SUCESSO")
	pause := ""
	fmt.Scan(&pause)
	misc.Limpatela() 

	return nil
}

func Read() (mc m.MoedaCripto, err error) {
	moedas := mc

	return moedas, nil
}

func Update(mc m.MoedaCripto, id string) error {
	return nil
}

func Delete(id string) error {
	return nil
}
