package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	ctx context.Context
	db  *mongo.Database
}

func InitDB(uri string) *DB {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err.Error())
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err.Error())
	}

	db := client.Database("mock_db")

	return &DB{ctx: ctx, db: db}
}
