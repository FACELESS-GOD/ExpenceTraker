package Utility

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseInstace *sql.DB

func InitialiseDatabaseConnection() {
	DB, err := sql.Open("mysql", "root:Admin@123@tcp/ExpenseTrackerAPI?charset=UTF8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DatabaseInstace = DB
}

func TerminateDatabaseConnection() {
	DatabaseInstace.Close()
}
