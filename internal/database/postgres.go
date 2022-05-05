package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)


type Postgres struct{
	HOST string
	USERNAME string
	PASSWORD string
	DATABASE string
	PORT string

}

func (post Postgres) GetConnection()(*gorm.DB, error){
	post.HOST = os.Getenv("PG_HOST")
	post.USERNAME = os.Getenv("PG_USER")
	post.PASSWORD = os.Getenv("PG_PASSWORD")
	post.DATABASE = os.Getenv("DB_NAME")
	post.PORT = os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", post.HOST, post.PORT, post.USERNAME, post.PASSWORD, post.DATABASE)
	connection, connectionError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if connectionError != nil{
		log.Println(connectionError)
	}

	log.Println("PG Connected")

	return connection, nil
}
