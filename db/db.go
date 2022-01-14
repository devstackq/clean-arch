package db

type StorageConnecter interface {
	InitDb() (interface{}, error)
	// SetConfig(user, password, host, port, dbName string)
	// InitMongoDb() (*mongo.Database, error)
	// InitPostgresDb() (*sql.DB, error)
}

type Config struct {
	user      string
	password  string
	host      string
	port      string
	tableName string
	dbName    string
}

func NewDbObject(typeDb, user, password, host, port, tableName, dbName string) StorageConnecter {
	if typeDb == "mongodb" {
		return &MongoDb{
			Config{user, password, host, port, tableName, dbName},
		}
	} else if typeDb == "postgresql" {
		return &PostgreSqlDb{
			Config{user, password, host, port, tableName, dbName},
		}
	}
	return nil
}
