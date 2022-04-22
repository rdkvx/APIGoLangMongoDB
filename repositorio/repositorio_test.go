package repositorio

import (
	"DesafioTecnico/mock"
	"DesafioTecnico/server/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
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

func TestCreate(t *testing.T) {
	mock := mock.Mock{}

	err := Create(&mock, *mockedCrypto)
	require.Nil(t, err)
}

func TestRead(t *testing.T) {
	mock := mock.Mock{}
	mock.SingleResult = mongo.SingleResult{}

	result, err := Read(&mock, "321646464")
	require.NotNil(t, err)
	require.NotNil(t, result)
}

func TestReadAll(t *testing.T) {}

func TestUpdate(t *testing.T) {}

func TestDelete(t *testing.T) {}
