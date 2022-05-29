package create_user_command

import (
	user_email "salvegame197/golangApi/src/domain/user/valueobject/email"
	user_first_name "salvegame197/golangApi/src/domain/user/valueobject/first_name"
	user_last_name "salvegame197/golangApi/src/domain/user/valueobject/last_name"
	user_password "salvegame197/golangApi/src/domain/user/valueobject/password"
	"time"
)

type CreateUserCommand struct {
	FirstName user_first_name.FirstName
	LastName  user_last_name.LastName
	Email     user_email.Email
	Password  user_password.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCreateUserCommand(
	FirstName string,
	LastName string,
	Email string,
	Password string,
) (*CreateUserCommand, error) {
	userFirstNameValueObject, _ := user_first_name.FromValue(FirstName)
	userLastNameValueObject, _ := user_last_name.FromValue(LastName)
	userEmailValueObject, err := user_email.FromValue(Email)
	if err != nil {
		return nil, err
	}
	userPasswordValueObject, _ := user_password.FromValue(Password)
	return &CreateUserCommand{
		FirstName: userFirstNameValueObject,
		LastName:  userLastNameValueObject,
		Email:     userEmailValueObject,
		Password:  userPasswordValueObject,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
