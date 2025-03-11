package Utility

import (
	"ExpenceTraker/Helper"
	"database/sql"
)

var DatabaseInstace *sql.DB

func InitialiseDatabaseConnection() {
	DB, err := sql.Open(Helper.DatabaseType, Helper.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}
	DatabaseInstace = DB
}

func TerminateDatabaseConnection() {
	DatabaseInstace.Close()
}
