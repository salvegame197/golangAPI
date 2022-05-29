package get_user_by_email

import (
	valueobject "salvegame197/golangApi/src/domain/user/valueobject/email"
)

type GetUserByEmailQuery struct {
	Email valueobject.Email
}

func NewGetUserByEmailQuery(
	Email string,
) *GetUserByEmailQuery {
	userEmailValueObject, _ := valueobject.FromValue(Email)
	return &GetUserByEmailQuery{
		Email: userEmailValueObject,
	}
}
