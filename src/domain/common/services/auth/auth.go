package auth_service

import (
	access_detail "salvegame197/golangApi/src/domain/common/valueobject/access"
	token_detail "salvegame197/golangApi/src/domain/common/valueobject/token"
)

type AuthInterface interface {
	CreateAuth(string, *token_detail.TokenDetail) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(authD *access_detail.AccessDetail) error
}
