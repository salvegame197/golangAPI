package create_user_command

import (
	"errors"
	security_service "salvegame197/golangApi/src/domain/common/services/security"
	create_user_service "salvegame197/golangApi/src/domain/user/service/create_user"
)

// Bus profile query bus
type Bus struct {
	service  *create_user_service.CreateUserService
	security security_service.Interface
}

// New create bus instance
func New(service *create_user_service.CreateUserService, security security_service.Interface) *Bus {
	return &Bus{service: service, security: security}
}

// Handle handle query
func (bus *Bus) Handle(command interface{}) error {
	switch command := command.(type) {
	case *CreateUserCommand:
		return bus.handleCreateUserCommand(command, bus.security)
	default:
		return errors.New("invalid command")
	}
}
