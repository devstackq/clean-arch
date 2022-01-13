package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageConnecter interface {
	InitDb() (interface{}, error)
	SetConfig(user, password, host, port, dbName string)
	// InitMongoDb() (*mongo.Database, error)
	// InitPostgresDb() (*sql.DB, error)
}

type Config struct {
	name         string
	password     string
	url          string
	port         string
	databaseName string
}

type DbFactory struct {
	user         string
	password     string
	host         string
	port         string
	databaseName string
}

type MongoDb struct {
	DbFactory
}
type PostgreSqlDb struct {
	DbFactory
}

func (df *DbFactory) SetConfig(user, password, host, port, dbName string) {
	df.user = user
	df.password = password
	df.host = host
	df.port = port
	df.databaseName = dbName
}

func GetDbFactory(dbName string) StorageConnecter {
	if dbName == "mongodb" {
		return &MongoDb{}
	} else if dbName == "postgresql" {
		return &PostgreSqlDb{}
	}
	return nil
}

func (p *PostgreSqlDb) InitDb() (interface{}, error) {
	db, err := sql.Open("postgres", p.DbFactory.host+"://"+p.DbFactory.port+p.DbFactory.user+"@"+p.DbFactory.password+p.DbFactory.databaseName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// if err = p.DbFactory.CreateTables(db); err != nil {
	// 	return nil, err
	// }
	return db, nil
}

func (m *MongoDb) InitDb() (interface{}, error) {
	log.Print("init mongo")
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
	// m.db = client.Database(m.Config.name)
	return client.Database(m.DbFactory.databaseName), nil
}
