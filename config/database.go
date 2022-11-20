package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
	Type       string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
}

func NewDatabaseConnection() Database {

	db := Database{
		Type:       Getenv("DB_CONNECTION", "mysql"),
		Host:       Getenv("DB_HOST", "127.0.0.1"),
		Port:       Getenv("DB_PORT", "8080"),
		Database:   Getenv("DB_DATABASE", ""),
		UserName:   Getenv("DB_USERNAME", "root"),
		Password:   Getenv("DB_PASSWORD", ""),
		Connection: &sql.DB{},
	}

	db.StartConnection()

	return db
}

func (d *Database) StartConnection() {

	db, err := sql.Open(d.Type, d.UserName+":"+d.Password+"@tcp("+d.Host+":"+d.Port+")/"+d.Database)

	if err != nil {
		log.Panic(err)
	}

	d.Connection = db
}

func (d *Database) Close() {
	defer d.Connection.Close()
}
