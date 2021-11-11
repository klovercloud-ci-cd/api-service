package logic

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"log"
	"time"
)

type jwtService struct {
	Jwt v1.Jwt
}

func (j jwtService) GenerateToken(duration int64, data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp":  time.Duration(duration) * time.Hour,
		"iat":  time.Now().Unix(),
		"data": data,
	}
	tokenString, err := token.SignedString(j.Jwt.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j jwtService) ValidateToken(tokenString string) (bool, *jwt.Token) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Jwt.PublicKey, nil
	})
	if err != nil {
		log.Print("[ERROR]: Token is invalid! ", err.Error())
		return false, nil
	}
	return true, token

}

func getPrivateKey() *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(config.PrivateKey))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(config.PublicKey))
	publicKeyImported, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return publicKeyImported
}

// NewJwtService returns Jwt type service
func NewJwtService() service.Jwt {
	return jwtService{
		Jwt: v1.Jwt{
			PrivateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		},
	}
}
