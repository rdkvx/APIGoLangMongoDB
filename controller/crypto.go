package controller

import (
	"DesafioTecnico/server/model"
	"DesafioTecnico/services"
	"DesafioTecnico/utils"
	"context"
	"errors"
	"fmt"
	"time"

	pb "DesafioTecnico/proto"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type CryptoServer struct {
	pb.UnimplementedCryptoServiceServer
	Coll *mongo.Collection
}

var Observer chan string

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
	utils.SkipLine()

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

	if server.Coll != nil {
		//Saving the new crypto on DB
		fmt.Println("CREATING A CRYPTO.......")
		err := services.Create(server.Coll, *newCrypto)
		if err != nil {
			return &pb.Cryptocurrency{}, err
		}
	}

	//Print the new crypto
	fmt.Println("CRYPTO CREATED SUCCESSFULLY")
	fmt.Println()
	fmt.Println("NAME      : ", newCrypto.Name)
	fmt.Println("SYMBOL    : ", newCrypto.Symbol)
	fmt.Println("VOTES     : ", newCrypto.Votes)
	fmt.Println("CREATED AT: ", newCrypto.CreatedAT)
	fmt.Println("UPDATED AT: ", newCrypto.UpdatedAT)
	utils.SkipLine()

	//If everythings gone right, return the new crypto object
	return &pb.Cryptocurrency{
		Id:        newCrypto.Id,
		Name:      newCrypto.Name,
		Symbol:    newCrypto.Symbol,
		Votes:     int32(newCrypto.Votes),
		Createdat: newCrypto.CreatedAT,
		Updatedat: newCrypto.UpdatedAT,
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
	fmt.Print("Symbol: : ", symbol)
	utils.SkipLine()

	//Check if the name and symbols are valid.
	if len(name) < 3 || name == "" {
		fmt.Println("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
		return &pb.Cryptocurrency{}, errors.New("INVALID NAME, MUST HAVE AT LEAST 3 CHARACTERS")
	}

	if len(symbol) < 3 || len(symbol) > 4 || symbol == "" {
		fmt.Println("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
		return &pb.Cryptocurrency{}, errors.New("INVALID SYMBOL, ONLY 3 OR 4 CHARACTERS ALLOWED")
	}
	if server.Coll != nil {

		//check if id is on DB
		res, err := services.Read(server.Coll, id)

		if err != nil {
			fmt.Println("FAILED TO UPDATE CRYPTO: INVALID ID")
			return &pb.Cryptocurrency{}, err
		} else {
			res.Name = name
			res.Symbol = symbol
			res.UpdatedAT = updatedat
		}

		//Update the data on DB
		err = services.Update(server.Coll, res)

		if err != nil {
			fmt.Println("FAILED TO UPDATE CRYPTO: ", err)
			return &pb.Cryptocurrency{}, err
		}

		//Print the new Crypto
		fmt.Println("CRYPTO UPDATED SUCCESSFULLY")
		utils.SkipLine()
		fmt.Println("NAME      : ", res.Name)
		fmt.Println("SYMBOL    : ", res.Symbol)
		fmt.Println("VOTES     : ", res.Votes)
		fmt.Println("CREATED AT: ", res.CreatedAT)
		fmt.Println("UPDATED AT: ", res.UpdatedAT)
		utils.SkipLine()

		//A Observer from datastream, if the ID that you defined on subscribe method
		//is the same beeing updated, the Observer will send a alert to the stream.
		Observer <- id

		//Return a object updated
		return &pb.Cryptocurrency{
			Id:        res.Id,
			Name:      res.Name,
			Symbol:    res.Symbol,
			Votes:     int32(res.Votes),
			Createdat: res.CreatedAT,
			Updatedat: res.UpdatedAT,
		}, nil
	}

	return &pb.Cryptocurrency{
		Id:        request.Id,
		Name:      request.Name,
		Symbol:    request.Symbol,
		Votes:     0,
		Createdat: "",
		Updatedat: "",
	}, nil
}

func (server *CryptoServer) Delete(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.EmptyResponse, error) {
	id := request.GetId()

	utils.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	utils.SkipLine()

	//check if ID is empty
	if id == "" {
		fmt.Println("FAILED TO DELETE A CRYPTO: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("FAILED TO DELETE A CRYPTO: EMPTY ID")
	}

	if server.Coll != nil {
		//check if id exists on DB
		_, err := services.Read(server.Coll, id)

		if err != nil {
			fmt.Println("FAILED TO DELETE CRYPTO: INVALID ID")
			return &pb.EmptyResponse{}, errors.New("FAILED TO DELETE A CRYPTO: INVALID ID")
		}

		//delete from DB
		err = services.Delete(server.Coll, id)
		if err != nil {
			fmt.Println("FAILED TO DELETE A CRYPTO: ", err)
		} else {
			fmt.Println("CRYPTO DELETED SUCCESSFULLY")
		}
	}

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Find(ctx context.Context, request *pb.FindRequest) (*pb.Cryptocurrency, error) {
	id := request.GetId()

	utils.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	utils.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO FIND CRYPTO: EMPTY ID")
		return &pb.Cryptocurrency{}, errors.New("FAILED TO FIND CRYPTO: EMPTY ID")
	}

	if server.Coll != nil {
		cryptoFound, err := services.Read(server.Coll, id)
		if err != nil {
			fmt.Println("FAILED TO DELETE CRYPTO: INVALID ID")
			return &pb.Cryptocurrency{}, errors.New("FAILED TO DELETE A CRYPTO: INVALID ID")
		}

		//Print the new Crypto
		fmt.Print("CRYPTO FOUND!")
		utils.SkipLine()
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
			Updatedat: cryptoFound.UpdatedAT,
		}, nil
	}

	return &pb.Cryptocurrency{
		Id:        id,
		Name:      "",
		Symbol:    "",
		Votes:     int32(0),
		Createdat: "",
		Updatedat: "",
	}, nil
}

func (server *CryptoServer) List(ctx context.Context, request *pb.ListCryptosRequest) (*pb.ListCryptosResponse, error) {
	sortParam := request.GetSortparam()
	ascending := request.GetAscending()

	utils.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Println("SORT PARAMETERS: ", sortParam)
	fmt.Print("ASCENDING        : ", ascending)
	utils.SkipLine()

	if server.Coll != nil {
		//get a collection with the parameters sended
		//ordered by: _id, name, symbol, voted, createdat, updatedat
		//and if its ordered by ascending order or descending
		//obs: if sortparams was sended empty, the default order will be by name
		cryptoList, err := services.ReadAll(server.Coll, sortParam, ascending)
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
				Updatedat: element.UpdatedAT,
			}
			cryptoPbList = append(cryptoPbList, newobj)
			utils.SkipLine()
			fmt.Println("ID         : ", element.Id)
			fmt.Println("NAME       : ", element.Name)
			fmt.Println("SYMBOL     : ", element.Symbol)
			fmt.Println("VOTES      : ", element.Votes)
			fmt.Println("CREATED AT : ", element.CreatedAT)
			fmt.Println("UPDATED AT : ", element.UpdatedAT)
		}
		utils.SkipLine()

		//return the list
		return &pb.ListCryptosResponse{
			Crypto: cryptoPbList,
		}, nil
	}

	return &pb.ListCryptosResponse{}, nil
}

func (server *CryptoServer) Upvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	id := request.GetId()

	utils.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	utils.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO VOTE: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("INVALID ID: EMPTY ID")
	}

	if server.Coll != nil {

		res, err := services.Read(server.Coll, id)
		if err != nil {
			fmt.Println("FAILED TO VOTE: CRYPTO NOT FOUND")
			return &pb.EmptyResponse{}, errors.New("FAILED TO VOTE: CRYPTO NOT FOUND")
		}

		//increment a vote
		res.Votes++

		err = services.Update(server.Coll, res)
		if err != nil {
			return &pb.EmptyResponse{}, nil
		}

		fmt.Println("VOTE REGISTERED SUCCESSFULLY!")
		fmt.Println("NAME: ", res.Name)
		fmt.Println("TOTAL VOTES: ", res.Votes)

		//A Observer from datastream, if the ID that you defined on subscribe method
		//is the same beeing updated, the Observer will send a alert to the stream.
		Observer <- id

	}

	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServer) Downvote(ctx context.Context, request *pb.VoteRequest) (*pb.EmptyResponse, error) {
	id := request.GetId()

	utils.SkipLine()
	fmt.Println("----------------------------------------")
	fmt.Print("PARAMETERS RECEIVED BY THE REQUEST \n\n")
	fmt.Print("ID: ", id)
	utils.SkipLine()

	if id == "" {
		fmt.Println("FAILED TO VOTE: EMPTY ID")
		return &pb.EmptyResponse{}, errors.New("INVALID ID: EMPTY ID")
	}

	if server.Coll != nil {
		res, err := services.Read(server.Coll, id)
		if err != nil {
			fmt.Println("FAILED TO VOTE: CRYPTO NOT FOUND")
			return &pb.EmptyResponse{}, errors.New("FAILED TO VOTE: CRYPTO NOT FOUND")
		}

		//decrement a vote
		if res.Votes == 0 {
			return &pb.EmptyResponse{}, errors.New("FAILED TO VOTE: VOTES ALREADY IN THE MINIMUM VALUE (0) CANNOT REGISTER NEGATIVE VOTES")
		}

		res.Votes--

		err = services.Update(server.Coll, res)
		if err != nil {
			return &pb.EmptyResponse{}, nil
		}

		fmt.Println("VOTE REGISTERED SUCCESSFULLY!")
		fmt.Println("NAME: ", res.Name)
		fmt.Println("TOTAL VOTES: ", res.Votes)

		//A Observer from datastream, if the ID that you defined on subscribe method
		//is the same beeing updated, the Observer will send a alert to the stream.
		Observer <- id
	}

	return &pb.EmptyResponse{}, nil
}

//Streaming watch method
func (server *CryptoServer) Subscribe(request *pb.SubscriptionRequest, stream pb.CryptoService_SubscribeServer) error {

	//keep the stream running
	for {
		cryptoUpdatedId := <-Observer //the ID is now beeing watched by the channel Observer

		if cryptoUpdatedId == request.Id {
			cryptoFound, err := services.Read(server.Coll, cryptoUpdatedId)
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
				Updatedat: cryptoFound.UpdatedAT,
			})
			if err != nil {
				fmt.Println("ERRO: ", err)
				return nil
			}
		}
	}
}
