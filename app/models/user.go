package models

import (
	"salvegame197/golangApi/config"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	gorm.Model
}

type UserServices interface {
	FindUserByID(a User, id int) (*User, error)
	FindUsers(a User) ([]*User, int, error)
	CreateUser(a User) error
	UpdateUser(a User, id int) (*User, error)
	DeleteUser(a User, id int) error
}

func FindUsers(a User) ([]*User, int, error) {
	var users []*User
	var count int
	err := config.DBmap.Model(&User{}).Where(a).Find(&users).Error
	return users, count, err
}
