package services

import (
	"DesafioTecnico/client/misc"
	m "DesafioTecnico/server/model"
	moeda_repositorio "DesafioTecnico/server/repositorio/moeda.repositorio"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CriarNovaCriptoMoedaAPI() error {
	mc := m.MoedaCripto{}
	
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

	err := moeda_repositorio.Create(mc)

	if err != nil {
		return err
	}

	return nil
}

func EditarCriptoMoedaAPI(id string, mc m.MoedaCripto) error {
	fmt.Print("\nNOVO NOME: ")
	fmt.Scan(&mc.Nome)

	fmt.Print("NOVO SIMBOLO: ")
	fmt.Scan(&mc.Simbolo)

	mc.UpdatedAT = time.Now().Format("02/01/2006 15:04:45")

	err := moeda_repositorio.Update(id, mc)

	if err != nil {
		return err
	}
	return nil
}

func DeletarCriptoMoedaAPI(id string) error {
	err := moeda_repositorio.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func UpVoteAPI(id string, mc m.MoedaCripto) {

}

func DownVoteAPI(id string, mc m.MoedaCripto) {

}

func ListarUmaCriptoAPI(id string) (m.MoedaCripto, error) {
	moeda, err := moeda_repositorio.Read(id)

	if err != nil {
		return moeda, err
	}
	return moeda, err
}

func ListarCriptoMoedasAPI() (lista []m.MoedaCripto) {

	return
}
