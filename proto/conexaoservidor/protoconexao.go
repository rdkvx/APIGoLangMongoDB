package conexaoservidor

import (
	pb "DesafioTecnico/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	enderecoDoServidor = "localhost:50051"
)

func ConectaServidor() (*grpc.ClientConn, pb.UpVoteServiceClient, context.Context, context.CancelFunc) {
	conexao, err := grpc.Dial(enderecoDoServidor, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error ao conectar ao servidor %v", err)
	}

	client := pb.NewUpVoteServiceClient(conexao)

	contexto, cancel := context.WithTimeout(context.Background(), time.Second*60)

	return conexao, client, contexto, cancel
}
