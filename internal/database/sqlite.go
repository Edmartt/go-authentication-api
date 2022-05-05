package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct{
	dbName string
}


func (con SQLite) GetConnection() (*gorm.DB, error){
	con.dbName = os.Getenv("DB_NAME")
	connection, connectionError := gorm.Open(sqlite.Open("./test.db"))
	
	if connectionError != nil{
		return nil, connectionError
	}

	return connection, nil
}


