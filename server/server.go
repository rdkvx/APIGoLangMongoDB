package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "DesafioTecnico/proto"
	"DesafioTecnico/repositorio"
	"DesafioTecnico/server/model"
	"time"
)

var (
	tipo  string = "tcp"
	porta string = ":50051"
)

type ProtoServer struct {
	pb.UnimplementedUpVoteServiceServer
}

func (server *ProtoServer) CriarNovaCriptoMoeda(ctx context.Context, request *pb.NovaCriptoRequest) (*pb.NovaCriptoResponse, error) {
	//Parametros recebidos do request
	id := request.GetId()
	nome := request.GetName()
	simbolo := request.GetSymbol()
	createdat := time.Now().Format("02/01/2006 15:04:45")

	//Validar valores recebidos. se nao forem validos, retornar error
	if len(nome) <= 3 {
		return &pb.NovaCriptoResponse{}, errors.New("NOME INVALIDO")
	}

	if len(simbolo) <= 1 {
		return &pb.NovaCriptoResponse{}, errors.New("SIMBOLO INVALIDO")
	}

	//Prosseguir usando valores válidos recebidos
	//Listar Valores recebidos no request
	fmt.Print("PARAMETROS RECEBIDOS PELO REQUEST \n\n")
	fmt.Println("Name: ", nome)
	fmt.Println("Symbol: ", simbolo)
	fmt.Println("Created At: ", createdat)

	//Instancia um novo model de cripto que será persistida no banco de dados
	novaCripto := &model.MoedaCripto{
		Id:        id,
		Nome:      nome,
		Simbolo:   simbolo,
		Voto:      0, //valor default é sempre zero para uma nova cripto criada
		CreatedAT: createdat,
		UpdatedAT: "",
	}

	//Persistir essa nova cripto no Banco de Dados e validar se deu certo ou teve erro
	fmt.Println("Gravando a nova cripto no Banco de Dados")
	err := repositorio.Create(*novaCripto)
	if err != nil {
		return &pb.NovaCriptoResponse{}, err
	}

	//Imprimir a nova cripto gravada
	fmt.Println(novaCripto)

	//Se deu certo, retornar um response com essa cripto gerada
	return &pb.NovaCriptoResponse{
		Id:        novaCripto.Id,
		Name:      novaCripto.Nome,
		Symbol:    novaCripto.Simbolo,
		Votes:     int32(novaCripto.Voto),
		Createdat: novaCripto.CreatedAT,
		Updatedat: novaCripto.UpdatedAT,
	}, nil
}

func main() {
	fmt.Println(tipo)
	fmt.Println(porta)

	lis, err := net.Listen(tipo, porta)
	if err != nil {
		fmt.Println("Failed to listen on port: ", porta, "error: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterUpVoteServiceServer(s, &ProtoServer{})
	fmt.Println("server listening at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
