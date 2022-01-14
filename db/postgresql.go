package db

import (
	"database/sql"
)

//set params outside
// func NewPostgresStorage(name, password, url, port, dbName string) *PostgreSql {}
// func (p *PostgreSql) InitPostgresDb() (*sql.DB, error) {}

type PostgreSqlDb struct {
	Config
}

func (p *PostgreSqlDb) InitDb() (interface{}, error) {
	db, err := sql.Open("postgres", p.Config.host+"://"+p.Config.port+p.Config.user+"@"+p.Config.password+p.Config.dbName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = createTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	linksTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS links  (
		id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
		url varchar(255) NOT NULL UNIQUE,
		short varchar(255) NOT NULL,
		createdtime timestamp
	)`)
	if err != nil {
		return err
	}
	linksTable.Exec()
	//etc tables create..

	return nil
}
