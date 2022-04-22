package mock

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mock struct {
	insertResult mongo.InsertOneResult
	SingleResult mongo.SingleResult
	cursor       mongo.Cursor
	updateResult mongo.UpdateResult
}

func (m *Mock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &m.insertResult, nil
}

func (m *Mock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return &m.SingleResult
}
func (m *Mock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	return &m.cursor, nil
}
func (m *Mock) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &m.updateResult, nil
}
func (m *Mock) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return &m.SingleResult
}
