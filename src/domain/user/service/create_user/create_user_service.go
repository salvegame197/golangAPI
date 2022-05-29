package create_user_service

import (
	"fmt"

	aggregate "salvegame197/golangApi/src/domain/user/aggregate"
	user_errors "salvegame197/golangApi/src/domain/user/errors"
	user_repository "salvegame197/golangApi/src/domain/user/repository"
	get_user_by_email_service "salvegame197/golangApi/src/domain/user/service/get_user_by_email"
)

// CreateUserService profile query bus
type CreateUserService struct {
	repository user_repository.Interface
	userFinder *get_user_by_email_service.GetUserByEmail
}

// New create bus instance
func New(repository user_repository.Interface, userFinder *get_user_by_email_service.GetUserByEmail) *CreateUserService {
	return &CreateUserService{repository: repository, userFinder: userFinder}
}

func (s *CreateUserService) Create(user *aggregate.User) error {
	userFound, _ := s.userFinder.GetUserByEmail(user.Email)
	if userFound != nil {
		return user_errors.AlreadyExists(fmt.Sprintf("User with email %s, already exists", user.Email.GetValue()))
	}
	return s.repository.Save(user)
}
