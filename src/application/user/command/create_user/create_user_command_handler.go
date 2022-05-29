package create_user_command

import (
	security_service "salvegame197/golangApi/src/domain/common/services/security"
	aggregate "salvegame197/golangApi/src/domain/user/aggregate"
	user_id "salvegame197/golangApi/src/domain/user/valueobject/id"
	user_password "salvegame197/golangApi/src/domain/user/valueobject/password"
)

func (bus *Bus) handleCreateUserCommand(
	command *CreateUserCommand,
	security security_service.Interface,
) error {
	userHashedPassword, err := security.Hash(command.Password.GetValue())
	if err != nil {
		return err
	}
	userHashedPasswordValueObject, _ := user_password.FromValue(string(userHashedPassword))
	user_id_value_object := user_id.NewUuid()

	return bus.service.Create(
		&aggregate.User{
			ID:        user_id_value_object,
			FirstName: command.FirstName,
			LastName:  command.LastName,
			Email:     command.Email,
			Password:  userHashedPasswordValueObject,
			CreatedAt: command.CreatedAt,
			UpdatedAt: command.UpdatedAt,
		},
	)
}
