package models

import (
	"gorm.io/gorm"
)

type User struct{
	Id string `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	gorm.Model
}
