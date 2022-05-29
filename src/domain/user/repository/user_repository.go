package user_repository

import (
	aggregate "salvegame197/golangApi/src/domain/user/aggregate"
	user_email "salvegame197/golangApi/src/domain/user/valueobject/email"
)

// Interface is interface of user repository
type Interface interface {
	Save(*aggregate.User) error
	GetUserByEmail(email user_email.Email) (*aggregate.User, error)
}
