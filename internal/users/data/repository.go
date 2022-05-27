package data

import (
	"github.com/Edmartt/go-authentication-api/internal/database"
	"github.com/Edmartt/go-authentication-api/internal/users/models"
	"log"
)

type UserRepository struct{
	DB database.IConnection
	user models.User
}


func NewRepository(){
	db := UserRepository{} 
	db.DB = database.SQLite{}
}


func (data UserRepository) Find(username string) (*models.User){
	connection, connError := data.DB.GetConnection()
	if connError != nil{
		log.Println("connection error")
	}
	
	connection.Where("username = ?", username).First(&data.user)
	
	if data.user.Id == ""{
		return &data.user
	}

	return &data.user
}

func (data UserRepository) Create(user models.User) string{
	connection, connError := data.DB.GetConnection()
	if connError != nil{
		log.Println("connection error")
	}
	connection.Create(&user)
	return user.Id
}

func (data UserRepository) Update(user models.User){
	connection, connError := data.DB.GetConnection()

	if connError != nil{
		log.Println("connection error")
	}
	connection.Save(&user)
}

func (data UserRepository) Delete(user models.User, id string){
	connection, connError := data.DB.GetConnection()

	if connError != nil{
		log.Println("connection error")
	}
	connection.Delete(&user, id)
}
