package logic

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
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

func (j jwtService) ValidateToken(tokenString string) (bool, *jwt.Token) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Jwt.PublicKey, nil
	})
	if err != nil {
		log.Print("[ERROR]: Token is invalid! ", err.Error())
		return false, nil
	}
	var tm time.Time
	switch iat := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}
	if time.Now().UTC().After(tm) {
		return false, nil
	}
	return true, token

}

func (j jwtService) ValidateTokenForInternalCall(tokenString string) (bool, *jwt.Token) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Jwt.PublicKeyForInternalCall, nil
	})
	if err != nil {
		return false, nil
	}

	var tm time.Time
	switch iat := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}
	if time.Now().UTC().After(tm) {
		return false, nil
	}
	return true, token
}

//func getPrivateKey() *rsa.PrivateKey {
//	block, _ := pem.Decode([]byte(config.PrivateKey))
//
//	privateKeyImported, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//	if err != nil {
//		log.Print(err.Error())
//		return nil
//	}
//
//	return privateKeyImported
//}

func getPublicKeyForInternalCall() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(config.PublicKeyForInternalCall))
	publicKeyImported, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return publicKeyImported
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
			//PrivateKey: getPrivateKey(),
			PublicKey:                getPublicKey(),
			PublicKeyForInternalCall: getPublicKeyForInternalCall(),
		},
	}
}
