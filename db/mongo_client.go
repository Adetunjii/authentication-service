package db

import (
	"context"
	"github.com/Adetunjii/lookapp/auth-service/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoClient struct {
	logger utils.AppLogger
	client *mongo.Client
}

type DB struct {
	db *mongo.Database
}

func NewMongoClient(logger utils.AppLogger, mongoURI string) *MongoClient {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Fatal("could not create mongo client: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal("could not connect to client: ", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal("could not ping client: ", err)
	}

	return &MongoClient{
		logger: logger,
		client: client,
	}
}

func (mc *MongoClient) ConnectToDatabase(name string) *DB {
	db := mc.client.Database(name)
	return &DB{
		db: db,
	}
}

func (mc *MongoClient) Close() error {
	err := mc.client.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

func (mdb *DB) GetCollection(collection string) *mongo.Collection {
	return mdb.db.Collection(collection)
}

func (mdb *DB) Name() string {
	return mdb.db.Name()
}
