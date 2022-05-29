package user_id

import (
	"fmt"
	custom_error "salvegame197/golangApi/src/domain/common/custom-error"

	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func FromValue(value string) (UserId, error) {
	err := assertValidValue(value)
	if err != nil {
		return UserId{}, err
	}
	return UserId{value: value}, err
}

func assertValidValue(value string) error {
	_, err := uuid.Parse(value)
	if err != nil {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid UserId %s, the value should a valid uuid", value))
	}
	return nil
}

func NewUuid() UserId {
	return UserId{value: uuid.New().String()}
}

func (n UserId) GetValue() string {
	return n.value
}
