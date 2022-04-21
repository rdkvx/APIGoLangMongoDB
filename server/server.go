package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sort"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"DesafioTecnico/client/misc"
	pb "DesafioTecnico/proto"
	repo "DesafioTecnico/repositorio"
	"DesafioTecnico/server/model"
	"time"
)

var (
	tipo string = "tcp"
	port string = ":50051"
)

type ProtoServer struct {
	pb.UnimplementedUpVoteServiceServer
}

func (server *ProtoServer) CreateACrypto(ctx context.Context, request *pb.RequestNewCrypto) (*pb.ResponseNewCrypto, error) {
	//Parametros recebidos do request
	id := uuid.NewString()
	name := request.GetName()
	symbol := request.GetSymbol()
	createdat := time.Now().Format("02/01/2006 15:04:45")

	//Validar valores recebidos. se nao forem validos, retornar error
	if len(name) <= 3 || name == "" {
		return &pb.ResponseNewCrypto{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) <= 2 || len(symbol) >= 5 {
		return &pb.ResponseNewCrypto{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}

	//Prosseguir usando valores válidos recebidos
	//Listar Valores recebidos no request
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("Name: ", name)
	fmt.Println("Symbol: ", symbol)
	fmt.Println("")

	//Instancia um novo model de cripto que será persistida no banco de dados
	newCrypto := &model.MoedaCripto{
		Id:        id,
		Name:      name,
		Symbol:    symbol,
		Votes:     0, //valor default é sempre zero para uma nova cripto criada
		CreatedAT: createdat,
		UpdatedAT: "",
	}

	//Persistir essa nova cripto no Banco de Dados e validar se deu certo ou teve erro
	fmt.Println("CREATING A CRYPTO.......")
	err := repo.Create(*newCrypto)
	if err != nil {
		return &pb.ResponseNewCrypto{}, err
	}

	//Print the new crypto
	fmt.Println("CRYPTO CREATED SUCCESSFULLY")
	fmt.Println()
	fmt.Println("NAME: ", newCrypto.Name)
	fmt.Println("SYMBOL: ", newCrypto.Symbol)
	fmt.Println("VOTES: ", newCrypto.Votes)
	fmt.Println("CREATED AT: ", newCrypto.CreatedAT)
	fmt.Println("UPDATED AT: ", newCrypto.UpdatedAT)
	misc.PulaLinha()

	//if everythings gone right, return the new crypto object
	return &pb.ResponseNewCrypto{
		ResponseCripto: &pb.CryptoCoin{
			Id:        newCrypto.Id,
			Name:      newCrypto.Name,
			Symbol:    newCrypto.Symbol,
			Votes:     int32(newCrypto.Votes),
			Createdat: newCrypto.CreatedAT,
			Updateat:  newCrypto.UpdatedAT,
		},
	}, nil
}

func (server *ProtoServer) EditACrypto(ctx context.Context, request *pb.RequestEditCrypto) (*pb.ResponseEditCrypto, error) {
	id := request.GetId()
	name := request.GetName()
	symbol := request.GetSymbol()
	updatedat := time.Now().Format("02/01/2006 15:04:45")

	//Validar valores recebidos. se nao forem validos, retornar error
	if len(name) < 3 || name == "" {
		return &pb.ResponseEditCrypto{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) < 3 || len(symbol) > 4 {
		return &pb.ResponseEditCrypto{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}

	//Go on with valid inputs
	//List the request values

	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("Name: ", name)
	fmt.Println("Symbol: ", symbol)
	fmt.Println("")

	res, err := repo.Read(id)

	if err != nil {
		fmt.Println("ERROR TRYING TO UPDATE CRYPTO: INVALID ID")
		return &pb.ResponseEditCrypto{}, nil
	} else {
		res.Name = name
		res.Symbol = symbol
		res.UpdatedAT = updatedat
	}

	err = repo.Update(res)
	if err != nil {
		log.Panic("ERROR TRYING TO EDITING CRYPTO ", err)
	}

	//Print the new Crypto
	fmt.Println("CRYPTO UPDATED SUCCESSFULLY")
	misc.PulaLinha()
	fmt.Println("NAME: ", res.Name)
	fmt.Println("SYMBOL: ", res.Symbol)
	fmt.Println("VOTES: ", res.Votes)
	fmt.Println("CREATED AT: ", res.CreatedAT)
	fmt.Println("UPDATED AT: ", res.UpdatedAT)
	misc.PulaLinha()

	return &pb.ResponseEditCrypto{
		ResponseCripto: &pb.CryptoCoin{
			Id:        res.Id,
			Name:      res.Name,
			Symbol:    res.Symbol,
			Votes:     int32(res.Votes),
			Createdat: res.CreatedAT,
			Updateat:  res.UpdatedAT,
		},
	}, nil
}

func (server *ProtoServer) DeleteACrypto(ctx context.Context, request *pb.RequestDeleteCrypto) (*pb.ResponseDeleteCrypto, error) {
	id := request.GetId()

	err := repo.Delete(id)
	misc.PulaLinha()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("ID: ", id)
	if err != nil {
		fmt.Println("FAILED TO DELETE A CRYPTO! CHECK THE ID PROVIDED")
	} else {
		fmt.Println("CRYPTO DELETED SUCCESSFULLY")
	}

	return &pb.ResponseDeleteCrypto{}, nil
}

func (server *ProtoServer) ListAllCryptosOrderedByVoteDesc(ctx context.Context, request *pb.RequestListAllCryptos) (*pb.ResponseListAllCryptos, error) {

	res, err := repo.ReadAll()

	fmt.Println("----------------------------------------")
	fmt.Println("")
	if err != nil {
		return &pb.ResponseListAllCryptos{}, errors.New("FAILED TO LIST CRYPTOS")
	}

	list := []*pb.CryptoCoin{}

	//get the data saved on DB and move it to a Proto List
	for _, element := range res {
		newobj := &pb.CryptoCoin{}
		newobj.Id = element.Id
		newobj.Name = element.Name
		newobj.Symbol = element.Symbol
		newobj.Votes = int32(element.Votes)
		newobj.Createdat = element.CreatedAT
		newobj.Updateat = element.CreatedAT
		list = append(list, newobj)
	}

	//Print the Crypto List
	fmt.Println("CRYPTO LIST")
	misc.PulaLinha()
	for _, element := range list {
		fmt.Println("NAME: ", element.Name)
		fmt.Println("SYMBOL: ", element.Symbol)
		fmt.Println("VOTES: ", element.Votes)
		fmt.Println("CREATED AT: ", element.Createdat)
		fmt.Println("UPDATED AT: ", element.Updateat)
		misc.PulaLinha()
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Votes > list[j].Votes
	})

	return &pb.ResponseListAllCryptos{
		ResponseCripto: list,
	}, nil
}

func (server *ProtoServer) ListAllCryptosOrderedByVoteAsc(ctx context.Context, request *pb.RequestListAllCryptos) (*pb.ResponseListAllCryptos, error) {

	res, err := repo.ReadAll()

	fmt.Println("----------------------------------------")
	fmt.Println("")
	if err != nil {
		return &pb.ResponseListAllCryptos{}, errors.New("FAILED TO LIST CRYPTOS")
	}

	list := []*pb.CryptoCoin{}

	//get the data saved on DB and move it to a Proto List
	for _, element := range res {
		newobj := &pb.CryptoCoin{}
		newobj.Id = element.Id
		newobj.Name = element.Name
		newobj.Symbol = element.Symbol
		newobj.Votes = int32(element.Votes)
		newobj.Createdat = element.CreatedAT
		newobj.Updateat = element.CreatedAT
		list = append(list, newobj)
	}

	fmt.Println("final list: ", list)

	//Print the Crypto Llist
	fmt.Println("CRYPTO LIST")
	misc.PulaLinha()
	for _, element := range list {
		fmt.Println("NAME: ", element.Name)
		fmt.Println("SYMBOL: ", element.Symbol)
		fmt.Println("VOTES: ", element.Votes)
		fmt.Println("CREATED AT: ", element.Createdat)
		fmt.Println("UPDATED AT: ", element.Updateat)
		misc.PulaLinha()
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Votes < list[j].Votes
	})

	return &pb.ResponseListAllCryptos{
		ResponseCripto: list,
	}, nil
}

func main() {
	lis, err := net.Listen(tipo, port)
	if err != nil {
		fmt.Println("FAILED TO LISTEN ON PORT: ", port, "ERROR: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterUpVoteServiceServer(s, &ProtoServer{})
	misc.CleanScreen()
	fmt.Println("SERVER LISTENING AT ", lis.Addr())
	misc.PulaLinha()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("FAILED TO SERVE: %v", err)
	}
}
