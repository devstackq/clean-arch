package db

import (
	"context"
	"database/sql"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Config
	// db *mongo.Database
}

func NewMongoStorage(name, password, url, port, dbName string) *Mongo {
	return &Mongo{
		Config: Config{name: name, password: password, url: url, port: port, databaseName: dbName},
	}
}

func (m *Mongo) InitPostgresDb() (*sql.DB, error) {
	return nil, nil
}
func (m *Mongo) InitMongoDb() (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(m.Config.url + m.Config.port))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	// m.db = client.Database(m.Config.name)

	return client.Database(m.Config.name), nil
}
