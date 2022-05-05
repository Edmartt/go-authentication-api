package data

import "github.com/Edmartt/go-authentication-api/internal/users/models"

type IUserRepository interface{
	find(id string) models.User
	create(user models.User)
	update(id string)
	delete(id string) string
}
