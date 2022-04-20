package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "DesafioTecnico/proto"

	"google.golang.org/grpc"
)

const (
	enderecoDoServidor = "localhost:50051"
)

func main() {
	//prepara a conexao para o server
	conexao, err := grpc.Dial(enderecoDoServidor, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error ao conectar ao servidor %v", err)
	}
	defer conexao.Close()

	//instancia um novo client a partir do protobuf usando essa coneao
	client := pb.NewUpVoteServiceClient(conexao)

	//define o contexto
	contexto, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//prepara os valores que serão enviados no "request"
	var nome = "teste cripto"
	var simbolo = "tsc"

	//chama a funcao do servidor remoto passando os parametros necessários
	resposta, err := client.CriarNovaCriptoMoeda(contexto, &pb.NovaCriptoRequest{Name: nome, Symbol: simbolo})
	if err != nil {
		log.Fatalf("error ao gravar uma nova cripto %v", err)
	}

	//se deu certo, exibir a nova cripto gerada
	fmt.Println("Nova cripto gerada:")
	fmt.Printf("\nId: %s", resposta.GetId())
	fmt.Printf("\nNome: %s", resposta.GetName())
	fmt.Printf("\nSimbolo: %s", resposta.GetSymbol())
	fmt.Printf("\nVotos: %d", resposta.GetVotes())
	fmt.Printf("\nData Criaçao: %s", resposta.GetCreatedat())
	fmt.Printf("\nData Atualizaçao: %s", resposta.GetUpdatedat())

}
