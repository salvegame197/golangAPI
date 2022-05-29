package get_user_by_email_service

import (
	"fmt"

	aggregate "salvegame197/golangApi/src/domain/user/aggregate"
	user_errors "salvegame197/golangApi/src/domain/user/errors"
	user_repository "salvegame197/golangApi/src/domain/user/repository"
	user_email "salvegame197/golangApi/src/domain/user/valueobject/email"
)

// GetUserByEmail query bus
type GetUserByEmail struct {
	repository user_repository.Interface
}

// New create bus instance
func New(repository user_repository.Interface) *GetUserByEmail {
	return &GetUserByEmail{repository: repository}
}

func (s *GetUserByEmail) GetUserByEmail(email user_email.Email) (*aggregate.User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		err = user_errors.NotFound(fmt.Sprintf("User with email %s, not found", email.GetValue()))
		return nil, err
	}
	return user, nil
}
