package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type StorageConnecter interface {
	// InitDb() (interface{}, error)
	InitMongoDb() (*mongo.Database, error)
	InitPostgresDb() (*sql.DB, error)
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

//confgi := COnfig{}; config.NewConfig() -> NewDb(config)

//Open/closed?
//type Home struct {window, door, nameRoom} -> TurnHeat(); TurnLight(),
//type HomeSmarter interface {TurnLight(), TurnHeat(), TurnCondicioner() }

//type Kitchen struct{wallpaper} TurnLight(); !change parent method; but override
//type BathRoom struct -> TurnHeat(){override own logic}; !change parent method; but override
