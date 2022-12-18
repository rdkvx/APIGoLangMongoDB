package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"DesafioTecnico/client/misc"
	"DesafioTecnico/controller"
	"DesafioTecnico/database"
	pb "DesafioTecnico/proto"
	"DesafioTecnico/services"
)

var (
	connType string = "tcp"
	port string = ":50051"
)

func main() {
	lis, err := net.Listen(connType, port)
	if err != nil {
		fmt.Println("FAILED TO LISTEN ON PORT: ", port, "ERROR: ", err)
	}

	controller.Observer = make(chan string)

	client, ctx, cancel, err := database.Connect()
	if err != nil {
		fmt.Println("ERROR TRYING TO CONNECT AT DB: ", err)
	}

	collection := client.Database(database.DB).Collection(database.COLLECTION)

	s := grpc.NewServer()

	server := controller.CryptoServer{
		Coll: collection,
	}
	
	pb.RegisterCryptoServiceServer(s, &server)
	misc.Clear()
	fmt.Println("SERVER LISTENING AT ", lis.Addr())
	misc.SkipLine()

	//if this is the first time running the server, this script will create 3 cryptos on DB
	services.CreateInitialData(collection)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("FAILED TO SERVE: %v", err)
	}

	database.Close(client, ctx, cancel)
}
