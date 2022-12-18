package services

import (
	"DesafioTecnico/mock"
	m "DesafioTecnico/server/model"
	"context"
	"errors"

	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func Create(col mock.ICollection, mc m.CryptoCurrency) error {
	_, err := col.InsertOne(context.Background(), mc)
	if err != nil {
		return err
	}

	return nil
}

func Read(col mock.ICollection, id string) (mc m.CryptoCurrency, err error) {
	filter := bson.M{"_id": id}
	//return a pointer that brings a object corresponding to the id passed by parameters.
	cur := col.FindOne(context.Background(), filter)
	if err != nil {
		return mc, err
	}

	var results = m.CryptoCurrency{}

	err = cur.Decode(&results)

	if err != nil {
		return mc, err
	}

	return results, nil
}

func ReadAll(col mock.ICollection, sortParam string, ascending bool) (obj []m.CryptoCurrency, err error) {
	var cur *mongo.Cursor
	findOptions := options.Find()

	switch sortParam {
	case "id", "name", "symbol", "votes", "createdat", "updatedat":
		if ascending { //If ascending = true, the return will be the default.
			findOptions.SetSort(bson.M{sortParam: 1})
			cur, err = col.Find(context.Background(), bson.M{}, findOptions)
			if err != nil {
				return obj, err
			}
		} else { //Sort by `sortParam` field descending (higher value first)
			findOptions.SetSort(bson.M{sortParam: -1})
			cur, err = col.Find(context.Background(), bson.M{}, findOptions)
			if err != nil {
				return obj, err
			}
		}
	default:
		//Return a list on default order, wich is ordered by name asc
		sortParam = "name"
		findOptions.SetSort(bson.M{sortParam: 1})
		cur, err = col.Find(context.Background(), bson.M{}, findOptions)
		if err != nil {
			return obj, err
		}
	}

	var results = []m.CryptoCurrency{}

	if err = cur.All(context.Background(), &results); err != nil {
		return obj, err
	}

	return results, err
}

func Update(col mock.ICollection, mc m.CryptoCurrency) error {
	filter := bson.M{"_id": bson.M{"$eq": mc.Id}}

	update := bson.M{
		"$set": bson.M{
			"name":      mc.Name,
			"symbol":    mc.Symbol,
			"votes":     mc.Votes,
			"createdat": mc.CreatedAT,
			"updatedat": mc.UpdatedAT,
		},
	}

	_, err := col.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(col mock.ICollection, id string) error {
	filter := bson.M{"_id": id}

	col.FindOneAndDelete(context.Background(), filter)

	return nil
}

func CreateInitialData(col mock.ICollection) error {

	sliceobj, err := ReadAll(col, "name", true)
	
	if err != nil {
		return errors.New("FAILED TO READ DB")
	}

	if len(sliceobj) == 0 {
		obj := m.CryptoCurrency{}
		obj.Id = uuid.NewString()
		obj.Name = "BITCOIN"
		obj.Symbol = "BTC"
		obj.CreatedAT = time.Now().Format("02/01/2006 15:04:45")
		Create(col, obj)

		obj.Id = uuid.NewString()
		obj.Name = "ETHEREUM"
		obj.Symbol = "ETH"
		obj.CreatedAT = time.Now().Format("02/01/2006 15:04:45")
		Create(col, obj)

		obj.Id = uuid.NewString()
		obj.Name = "KLEVER"
		obj.Symbol = "KLV"
		obj.CreatedAT = time.Now().Format("02/01/2006 15:04:45")
		Create(col, obj)

		fmt.Println("INITIAL DATA CREATED")
	}

	return nil
}
