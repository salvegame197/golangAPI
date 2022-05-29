package security_service

import (
	security_service "salvegame197/golangApi/src/domain/common/services/security"

	"golang.org/x/crypto/bcrypt"
)

type Security struct {
}

func New() security_service.Interface {
	return &Security{}
}

func (r *Security) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (r *Security) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
