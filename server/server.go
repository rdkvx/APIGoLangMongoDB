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
	//Request parameters
	id := uuid.NewString()
	name := request.GetName()
	symbol := request.GetSymbol()
	createdat := time.Now().Format("02/01/2006 15:04:45")

	//Print the values received by the request
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("Nam   : ", name)
	fmt.Print("Symbol: ", symbol)
	misc.SkipLine()

	//Check if the name and symbols are valid.
	if len(name) <= 3 || name == "" {
		fmt.Println("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
		return &pb.Cryptocurrency{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) <= 2 || len(symbol) >= 5 || symbol == "" {
		fmt.Println("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
		return &pb.Cryptocurrency{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}

	//Instance a object of cryptoCurrency
	newCrypto := &model.CryptoCurrency{
		Id:        id,
		Name:      name,
		Symbol:    symbol,
		Votes:     0,
		CreatedAT: createdat,
		UpdatedAT: "",
	}

	//Saving the new crypto on DB
	fmt.Println("CREATING A CRYPTO.......")
	err := repo.Create(*newCrypto)
	if err != nil {
		return &pb.Cryptocurrency{}, err
	}

	//Print the new crypto
	fmt.Println("CRYPTO CREATED SUCCESSFULLY")
	fmt.Println()
	fmt.Println("NAME      : ", newCrypto.Name)
	fmt.Println("SYMBOL    : ", newCrypto.Symbol)
	fmt.Println("VOTES     : ", newCrypto.Votes)
	fmt.Println("CREATED AT: ", newCrypto.CreatedAT)
	fmt.Println("UPDATED AT: ", newCrypto.UpdatedAT)
	misc.SkipLine()

	//If everythings gone right, return the new crypto object
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

	//List the request values

	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("Id    : ", id)
	fmt.Println("Name  : ", name)
	fmt.Print("Symbol: ", symbol)
	misc.SkipLine()

	//Check if the name and symbols are valid.
	if len(name) < 3 || name == "" {
		fmt.Println("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
		return &pb.Cryptocurrency{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) < 3 || len(symbol) > 4 || symbol == "" {
		fmt.Println("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
		return &pb.Cryptocurrency{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}

	//check if id is on DB
	res, err := repo.Read(id)

	if err != nil {
		fmt.Println("FAILED TO UPDATE CRYPTO: INVALID ID")
		return &pb.Cryptocurrency{}, nil
	} else {
		res.Name = name
		res.Symbol = symbol
		res.UpdatedAT = updatedat
	}

	//Update the data on DB
	err = repo.Update(res)

	if err != nil {
		fmt.Println("FAILED TO UPDATE CRYPTO: ", err)
		return &pb.Cryptocurrency{}, err
	}

	//Print the new Crypto
	fmt.Println("CRYPTO UPDATED SUCCESSFULLY")
	misc.SkipLine()
	fmt.Println("NAME      : ", res.Name)
	fmt.Println("SYMBOL    : ", res.Symbol)
	fmt.Println("VOTES     : ", res.Votes)
	fmt.Println("CREATED AT: ", res.CreatedAT)
	fmt.Println("UPDATED AT: ", res.UpdatedAT)
	misc.SkipLine()

	//A observer from datastream, if the ID that you defined on subscribe method
	//is the same beeing updated, the observer will send a alert to the stream.
	observer <- id

	//Return a object updated
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
	id := request.GetId()

	misc.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	misc.SkipLine()

	//check if ID is empty
	if id == "" {
		fmt.Println("FAILED TO DELETE A CRYPTO: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("FAILED TO DELETE A CRYPTO: EMPTY ID")
	}

	//check if id exists on DB
	_, err := repo.Read(id)

	if err != nil {
		fmt.Println("FAILED TO DELETE CRYPTO: INVALID ID")
		return &pb.EmptyResponse{}, errors.New("FAILED TO DELETE A CRYPTO: INVALID ID")
	}
	//delete from DB
	err = repo.Delete(id)
	if err != nil {
		fmt.Println("FAILED TO DELETE A CRYPTO: ", err)
	} else {
		fmt.Println("CRYPTO DELETED SUCCESSFULLY")
	}

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Find(ctx context.Context, request *pb.FindRequest) (*pb.Cryptocurrency, error) {
	id := request.GetId()

	misc.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	misc.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO FIND CRYPTO: EMPTY ID")
		return &pb.Cryptocurrency{}, errors.New("FAILED TO FIND CRYPTO: EMPTY ID")
	}

	cryptoFound, err := repo.Read(id)
	if err != nil {
		fmt.Println("FAILED TO DELETE CRYPTO: INVALID ID")
		return &pb.Cryptocurrency{}, errors.New("FAILED TO DELETE A CRYPTO: INVALID ID")
	}

	//Print the new Crypto
	fmt.Print("CRYPTO FOUND!")
	misc.SkipLine()
	fmt.Println("NAME      : ", cryptoFound.Name)
	fmt.Println("SYMBOL    : ", cryptoFound.Symbol)
	fmt.Println("VOTES     : ", cryptoFound.Votes)
	fmt.Println("CREATED AT: ", cryptoFound.CreatedAT)
	fmt.Println("UPDATED AT: ", cryptoFound.UpdatedAT)
	

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
	sortParam := request.GetSortparam()
	ascending := request.GetAscending()

	misc.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("SORT PARAMETERS: ", sortParam)
	fmt.Print("ASCENDING      : ", ascending)
	misc.SkipLine()

	//get a collection with the parameters sended
	//ordered by: _id, name, symbol, voted, createdat, updatedat
	//and if its ordered by ascending order or descending
	//obs: if sortparams was sended empty, the default order will be by name
	cryptoList, err := repo.ReadAll(sortParam, ascending)
	if err != nil {
		fmt.Println("FAILED TO LIST CRYPTOS")
		return &pb.ListCryptosResponse{}, errors.New("FAILED TO LIST CRYPTOS")
	}

	//a list to receive the data from DB
	cryptoPbList := []*pb.Cryptocurrency{}

	//filling the list with the data from DB and displaying at server console.
	if sortParam == "" {
		fmt.Println("CRYPTO LIST ORDERED BY: NAME")
	} else {
		fmt.Println("CRYPTO LIST ORDERED BY: ", sortParam)
	}

	for _, element := range cryptoList {
		newobj := &pb.Cryptocurrency{
			Id:        element.Id,
			Name:      element.Name,
			Symbol:    element.Symbol,
			Votes:     int32(element.Votes),
			Createdat: element.CreatedAT,
			Updateat:  element.CreatedAT,
		}
		cryptoPbList = append(cryptoPbList, newobj)
		misc.SkipLine()
		fmt.Println("ID         : ", element.Id)
		fmt.Println("NAME       : ", element.Name)
		fmt.Println("SYMBOL     : ", element.Symbol)
		fmt.Println("VOTES      : ", element.Votes)
		fmt.Println("CREATED AT : ", element.CreatedAT)
		fmt.Println("UPDATED AT : ", element.UpdatedAT)
	}
	misc.SkipLine()

	//return the list
	return &pb.ListCryptosResponse{
		Crypto: cryptoPbList,
	}, nil
}

func (server *CryptoServer) Upvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	id := request.GetId()

	misc.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	misc.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO VOTE: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("INVALID ID: EMPTY ID")
	}

	res, err := repo.Read(id)
	if err != nil {
		fmt.Println("FAILED TO VOTE: CRYPTO NOT FOUND")
		return &pb.EmptyResponse{}, errors.New("FAILED TO VOTE: CRYPTO NOT FOUND")
	}

	//increment a vote
	res.Votes++

	err = repo.Update(res)
	if err != nil {
		return &pb.EmptyResponse{}, nil
	}

	fmt.Println("VOTE REGISTERED SUCCESSFULLY!")
	fmt.Println("NAME: ", res.Name)
	fmt.Println("TOTAL VOTES: ", res.Votes)

	//A observer from datastream, if the ID that you defined on subscribe method
	//is the same beeing updated, the observer will send a alert to the stream.
	observer <- id

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Downvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	id := request.GetId()

	misc.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	misc.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO VOTE: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("INVALID ID: EMPTY ID")
	}

	res, err := repo.Read(id)
	if err != nil {
		fmt.Println("FAILED TO VOTE: CRYPTO NOT FOUND")
		return &pb.EmptyResponse{}, errors.New("FAILED TO VOTE: CRYPTO NOT FOUND")
	}

	//decrement a vote
	res.Votes--

	err = repo.Update(res)
	if err != nil {
		return &pb.EmptyResponse{}, nil
	}

	fmt.Println("VOTE REGISTERED SUCCESSFULLY!")
	fmt.Println("NAME: ", res.Name)
	fmt.Println("TOTAL VOTES: ", res.Votes)

	//A observer from datastream, if the ID that you defined on subscribe method
	//is the same beeing updated, the observer will send a alert to the stream.
	observer <- id

	return &pb.EmptyResponse{}, nil
}

//Streaming watch method
func (server *CryptoServer) Subscribe(request *pb.SubscriptionRequest, stream pb.CryptoService_SubscribeServer) error {
	keepActive := true

	//keep the stream running
	for keepActive {
		cryptoUpdatedId := <-observer //the ID is now beeing watched by the channel observer

		if cryptoUpdatedId == request.Id {
			cryptoFound, err := repo.Read(cryptoUpdatedId)
			if err != nil {
				fmt.Println("UNABLE TO STREAM CHANGES IN CRYPTO FROM ID: ", cryptoUpdatedId)
				return nil
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
				fmt.Println("ERRO: ",err)
				keepActive = false
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
	misc.Clear()
	fmt.Println("SERVER LISTENING AT ", lis.Addr())
	misc.SkipLine()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("FAILED TO SERVE: %v", err)
	}
}
