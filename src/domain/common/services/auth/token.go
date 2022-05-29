package auth_service

import (
	"net/http"
	access_detail "salvegame197/golangApi/src/domain/common/valueobject/access"
	token_detail "salvegame197/golangApi/src/domain/common/valueobject/token"
)

type TokenInterface interface {
	CreateToken(userid string) (*token_detail.TokenDetail, error)
	ExtractTokenMetadata(*http.Request) (*access_detail.AccessDetail, error)
}
