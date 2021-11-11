package service

import (
	"github.com/dgrijalva/jwt-go"
)

// Jwt Jwt operations.
type Jwt interface {
	GenerateToken(duration int64, data interface{}) (string, error)
	ValidateToken(tokenString string) (bool, *jwt.Token)
}
