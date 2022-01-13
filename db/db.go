package db

type StorageConnecter interface {
	InitDb() (interface{}, error)
	SetConfig(user, password, host, port, dbName string)
	// InitMongoDb() (*mongo.Database, error)
	// InitPostgresDb() (*sql.DB, error)
}

type DbFactory struct {
	user         string
	password     string
	host         string
	port         string
	databaseName string
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
