package db

import (
	"context"
	"database/sql"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageConnecter interface {
	InitDb() (interface{}, error)

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

//base struct method, can override (Postgresql)NewConfig()
func (c *Config) NewConfig(name, password, url, port, databaseName string) Config {
	return Config{
		name: name, password: password,
		port: port, databaseName: databaseName,
		url: url,
	}
}

type DbFactory struct {
	typeDb       string
	user         string
	password     string
	host         string
	port         string
	databaseName string
}

func NewDbFactory(typeDb, user, password, host, port, database string) StorageConnecter {
	return &DbFactory{typeDb, user, password, host, port, database}
}

func (f *DbFactory) InitDb() (interface{}, error) {
	if f.typeDb == "mongo" {
		return f.InitMongoDb()
	} else if f.typeDb == "postgres" {
		return f.InitPSql()
	}
	return nil, nil
}
func (f *DbFactory) InitPSql() (interface{}, error) {
	db, err := sql.Open("postgres", f.host+"://"+f.port+f.user+"@"+f.password+f.databaseName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// if err = f.CreateTables(db); err != nil {
	// 	return nil, err
	// }
	return db, nil
}

func (f *DbFactory) InitMongoDb() (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(f.host + f.port))
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
	return client.Database(f.databaseName), nil
}
