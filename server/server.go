package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

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

type CryptoServer struct {
	pb.UnimplementedCryptoServiceServer
}

var observer chan string

func (server *CryptoServer) Create(ctx context.Context, request *pb.NewCryptoRequest) (*pb.Cryptocurrency, error) {
	//Parametros recebidos do request
	id := uuid.NewString()
	name := request.GetName()
	symbol := request.GetSymbol()
	createdat := time.Now().Format("02/01/2006 15:04:45")

	//Validar valores recebidos. se nao forem validos, retornar error
	if len(name) <= 3 || name == "" {
		return &pb.Cryptocurrency{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) <= 2 || len(symbol) >= 5 {
		return &pb.Cryptocurrency{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}

	//Prosseguir usando valores válidos recebidos
	//Listar Valores recebidos no request
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("Name: ", name)
	fmt.Println("Symbol: ", symbol)
	fmt.Println("")

	//Instancia um novo model de cripto que será persistida no banco de dados
	newCrypto := &model.CryptoCurrency{
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
		return &pb.Cryptocurrency{}, err
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
	return &pb.Cryptocurrency{
		Id:        newCrypto.Id,
		Name:      newCrypto.Name,
		Symbol:    newCrypto.Symbol,
		Votes:     int32(newCrypto.Votes),
		Createdat: newCrypto.CreatedAT,
		Updateat:  newCrypto.UpdatedAT,
	}, nil
}

func (server *CryptoServer) Edit(ctx context.Context, request *pb.EditCryptoRequest) (*pb.Cryptocurrency, error) {
	id := request.GetId()
	name := request.GetName()
	symbol := request.GetSymbol()
	updatedat := time.Now().Format("02/01/2006 15:04:45")

	//Validar valores recebidos. se nao forem validos, retornar error
	if len(name) < 3 || name == "" {
		return &pb.Cryptocurrency{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) < 3 || len(symbol) > 4 {
		return &pb.Cryptocurrency{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
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
		return &pb.Cryptocurrency{}, nil
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

	observer <- id

	return &pb.Cryptocurrency{
		Id:        res.Id,
		Name:      res.Name,
		Symbol:    res.Symbol,
		Votes:     int32(res.Votes),
		Createdat: res.CreatedAT,
		Updateat:  res.UpdatedAT,
	}, nil
}

func (server *CryptoServer) Delete(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.EmptyResponse, error) {
	//validar se recebeu um id mesmo
	id := request.GetId()

	misc.PulaLinha()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("ID: ", id)

	//deletar no banco
	err := repo.Delete(id)
	if err != nil {
		fmt.Println("FAILED TO DELETE A CRYPTO! CHECK THE ID PROVIDED")
	} else {
		fmt.Println("CRYPTO DELETED SUCCESSFULLY")
	}

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Find(ctx context.Context, request *pb.FindRequest) (*pb.Cryptocurrency, error) {
	id := request.GetId()

	cryptoFound, err := repo.Read(id)
	if err != nil {
		return &pb.Cryptocurrency{}, err
	}

	return &pb.Cryptocurrency{
		Id:        cryptoFound.Id,
		Name:      cryptoFound.Name,
		Symbol:    cryptoFound.Symbol,
		Votes:     int32(cryptoFound.Votes),
		Createdat: cryptoFound.CreatedAT,
		Updateat:  cryptoFound.UpdatedAT,
	}, nil
}

func (server *CryptoServer) List(ctx context.Context, request *pb.ListCryptosRequest) (*pb.ListCryptosResponse, error) {
	//Pegar valores recebidos no request
	sortParam := request.GetSortparam()
	ascending := request.GetAscending()

	//pega a collection de model crypto do mongodb
	cryptoList, err := repo.ReadAll(sortParam, ascending)
	if err != nil {
		return &pb.ListCryptosResponse{}, errors.New("FAILED TO LIST CRYPTOS")
	}

	//cria uma list de crypto protobuf pra ser retornada pelo método
	cryptoPbList := []*pb.Cryptocurrency{}

	//itera na lista do mongo para converter o modelo do go pro modelo do protobuf
	for _, element := range cryptoList {
		newobj := &pb.Cryptocurrency{
			Id:        element.Id,
			Name:      element.Name,
			Symbol:    element.Symbol,
			Votes:     int32(element.Votes),
			Createdat: element.CreatedAT,
			Updateat:  element.CreatedAT,
		}

		//concatena o elemento iterado na lista de retorno
		cryptoPbList = append(cryptoPbList, newobj)
	}

	//retorna o resultado
	return &pb.ListCryptosResponse{
		Crypto: cryptoPbList,
	}, nil
}

func (server *CryptoServer) Upvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	//pega o id pra inscrementar o voto
	cryptoId := request.GetId()
	if cryptoId == "" {
		return &pb.EmptyResponse{}, errors.New("invalid Id")
	}

	res, err := repo.Read(cryptoId)
	if err != nil {
		return &pb.EmptyResponse{}, errors.New("crypto not found")
	}

	//incrmenta o voto
	res.Votes = res.Votes + 1

	err = repo.Update(res)
	if err != nil {
		return &pb.EmptyResponse{}, nil
	}

	observer <- cryptoId

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Downvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	//pega o id pra inscrementar o voto
	cryptoId := request.GetId()
	if cryptoId == "" {
		return &pb.EmptyResponse{}, errors.New("invalid Id")
	}

	res, err := repo.Read(cryptoId)
	if err != nil {
		return &pb.EmptyResponse{}, errors.New("crypto not found")
	}

	//incrmenta o voto
	res.Votes = res.Votes - 1

	err = repo.Update(res)
	if err != nil {
		return &pb.EmptyResponse{}, nil
	}

	observer <- cryptoId

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Subscribe(request *pb.SubscriptionRequest, stream pb.CryptoService_SubscribeServer) error {
	keepActive := true

	for keepActive {
		cryptoUpdatedId := <-observer

		if cryptoUpdatedId == request.Id {
			cryptoFound, err := repo.Read(cryptoUpdatedId)
			if err != nil {
				fmt.Println("Unable to stream changes in crypto from Id: ", cryptoUpdatedId)
			}

			err = stream.Send(&pb.Cryptocurrency{
				Id:        cryptoFound.Id,
				Name:      cryptoFound.Name,
				Symbol:    cryptoFound.Symbol,
				Votes:     int32(cryptoFound.Votes),
				Createdat: cryptoFound.CreatedAT,
				Updateat:  cryptoFound.UpdatedAT,
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen(tipo, port)
	if err != nil {
		fmt.Println("FAILED TO LISTEN ON PORT: ", port, "ERROR: ", err)
	}

	observer = make(chan string)

	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, &CryptoServer{})
	misc.CleanScreen()
	fmt.Println("SERVER LISTENING AT ", lis.Addr())
	misc.PulaLinha()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("FAILED TO SERVE: %v", err)
	}
}
