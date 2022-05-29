package data

import "github.com/Edmartt/go-authentication-api/internal/users/models"

var RepoAccessInterface IUserRepository

type IUserRepository interface{
	Find(id string) models.User
	Create(user models.User) string
}
