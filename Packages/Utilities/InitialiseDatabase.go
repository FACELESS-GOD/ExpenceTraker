package Utility

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseInstace *sql.DB

func InitialiseDatabaseConnection() {
	DB, err := sql.Open()
	if err != nil {
		panic(err)
	}
	DatabaseInstace = DB
}

func TerminateDatabaseConnection() {
	DatabaseInstace.Close()
}
