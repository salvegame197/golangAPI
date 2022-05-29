package get_user_by_email

import (
	"errors"

	user_aggregate "salvegame197/golangApi/src/domain/user/aggregate"
	get_user_by_email_service "salvegame197/golangApi/src/domain/user/service/get_user_by_email"
)

// Bus profile query bus
type Bus struct {
	service *get_user_by_email_service.GetUserByEmail
}

// New create bus instance
func New(service *get_user_by_email_service.GetUserByEmail) *Bus {
	return &Bus{service: service}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*user_aggregate.User, error) {
	switch query := query.(type) {
	case *GetUserByEmailQuery:
		return bus.handleFindUserByEmailQuery(query)
	default:
		return nil, errors.New("invalid query")
	}
}
