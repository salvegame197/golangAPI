package get_user_by_email

import aggregate "salvegame197/golangApi/src/domain/user/aggregate"

func (bus *Bus) handleFindUserByEmailQuery(
	query *GetUserByEmailQuery,
) (*aggregate.User, error) {
	return bus.service.GetUserByEmail(
		query.Email,
	)
}
