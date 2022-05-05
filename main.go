package main

import (
	"log"
	"os"

	server "github.com/Edmartt/go-authentication-api/internal"
	"github.com/Edmartt/go-authentication-api/internal/database"

	"github.com/joho/godotenv"
)

func main(){
	envError := godotenv.Load()

	if envError != nil{
		log.Fatal("Failed loading .env file")
	}

	migration := database.Migrations{DB: database.Postgres{}}
	migration.MakeMigrations()

	server := server.HttpServer{}
	server.SetServer()
	server.StartServer(os.Getenv("HTTP_PORT"))
}
