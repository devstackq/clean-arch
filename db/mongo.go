package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	DbFactory
}

func (m *MongoDb) InitDb() (interface{}, error) {
	log.Print(m.DbFactory.host, m.DbFactory.port)
	client, err := mongo.NewClient(options.Client().ApplyURI(m.DbFactory.host + m.DbFactory.port))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	log.Print("init mongo")

	// m.db = client.Database(m.Config.name)
	return client.Database(m.DbFactory.databaseName), nil
}
