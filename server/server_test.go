package main

import (
	pb "DesafioTecnico/proto"
	"DesafioTecnico/server/model"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var mockedCrypto = &model.CryptoCurrency{
	Id:        uuid.NewString(),
	Name:      "mocked crypto",
	Symbol:    "mkd",
	Votes:     35,
	CreatedAT: time.Now().Format("02/01/2006 15:04:45"),
	UpdatedAT: "",
}

var mockedCryptoEmpty = &model.CryptoCurrency{
	Id:        "",
	Name:      "",
	Symbol:    "",
	Votes:     0,
	CreatedAT: "",
	UpdatedAT: "",
}

var pbMockedCrypto = &pb.NewCryptoRequest{
	Name:   "",
	Symbol: mockedCrypto.Symbol,
}

var pbMockedCryptoEmpty = &pb.Cryptocurrency{
	Id:        "",
	Name:      "",
	Symbol:    "",
	Votes:     0,
	Createdat: "",
	Updateat:  "",
}

var pbMockedEditCrypto = &pb.EditCryptoRequest{
	Id:     "123456",
	Name:   "Edited Crypto",
	Symbol: "EDC",
}

var pbMockedDeleteCrypto = &pb.DeleteCryptoRequest{
	Id: "123",
}

var pbMockedFindCrypto = &pb.FindRequest{
	Id: "123",
}

var pbMockedVoteCrypto = &pb.VoteRequest{
	Id: "123",
}

func TestCreate(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	//success success
	result, err := server.Create(ctx, &pb.NewCryptoRequest{
		Name:   mockedCrypto.Name,
		Symbol: mockedCrypto.Symbol,
	})
	require.Nil(t, err)
	require.NotNil(t, result)
	//Sent values
	require.Equal(t, mockedCrypto.Name, result.Name)
	require.Equal(t, mockedCrypto.Symbol, result.Symbol)
	//values gerenated inside Create method from server
	require.NotEmpty(t, result.Id)
	require.NotEmpty(t, result.Createdat)
	require.Empty(t, result.Updateat)

	//fail case
	result, err = server.Create(ctx, &pb.NewCryptoRequest{
		Name:   "",
		Symbol: mockedCrypto.Symbol,
	})
	require.NotNil(t, err)
	require.Equal(t, pbMockedCryptoEmpty, result)

	defer cancel()
}

func TestEdit(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.Edit(ctx, &pb.EditCryptoRequest{
		Id:     pbMockedEditCrypto.Id,
		Name:   pbMockedEditCrypto.Name,
		Symbol: pbMockedEditCrypto.Symbol,
	})

	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, pbMockedEditCrypto.Id, result.Id)
	require.Equal(t, pbMockedEditCrypto.Name, result.Name)
	require.Equal(t, pbMockedEditCrypto.Symbol, result.Symbol)
	require.Equal(t, int32(0), result.Votes)
	require.Equal(t, "", result.Createdat)
	require.Equal(t, "", result.Updateat)

	defer cancel()
}

func TestDelete(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.Delete(ctx, &pb.DeleteCryptoRequest{
		Id: pbMockedDeleteCrypto.Id,
	})

	require.Nil(t, err)
	require.NotNil(t, result)

	defer cancel()
}

func TestFind(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.Find(ctx, &pb.FindRequest{
		Id: pbMockedFindCrypto.Id,
	})

	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, pbMockedFindCrypto.Id, result.Id)

	defer cancel()
}

func TestList(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.List(ctx, &pb.ListCryptosRequest{
		Sortparam: "votes",
		Ascending: true,
	})

	require.Nil(t, err)
	require.NotNil(t, result)

	defer cancel()
}

func TestUpVote(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.Upvote(ctx, &pb.VoteRequest{
		Id: pbMockedVoteCrypto.Id,
	})

	require.Nil(t, err)
	require.NotNil(t, result)

	defer cancel()
}

func TestDownVote(t *testing.T) {
	server := CryptoServer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	result, err := server.Downvote(ctx, &pb.VoteRequest{
		Id: pbMockedVoteCrypto.Id,
	})

	require.Nil(t, err)
	require.NotNil(t, result)

	defer cancel()
}

func TestSubscribe(t *testing.T) {

}
