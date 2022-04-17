package services

import (
	m "DesafioTecnico/server/model"
	moeda_repositorio "DesafioTecnico/server/repositorio/moeda.repositorio"
)

func CriarNovaCriptoMoedaAPI(mc m.MoedaCripto) error{
	err := moeda_repositorio.Create(mc)

	if err != nil{
		return err
	}

	return nil
}

func EditarCriptoMoedaAPI(id string, mc m.MoedaCripto) error{
	err := moeda_repositorio.Update(mc, id)

	if err != nil {
		return err
	}
	return nil
}

func DeletarCriptoMoedaAPI(id string) error{
	err := moeda_repositorio.Delete(id)

	if err != nil{
		return err
	}
	return nil
}

func UpVoteAPI(id string, mc m.MoedaCripto) {
	
}

func DownVoteAPI(id string, mc m.MoedaCripto) {

}

func ListarUmaCriptoAPI() (m.MoedaCripto, error){
	moedas, err := moeda_repositorio.Read()

	if err != nil{
		return moedas, err
	}
	return moedas, err
}

func ListarCriptoMoedasAPI() (lista []m.MoedaCripto) {

	return
}
