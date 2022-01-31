package v1

import (
	"crypto/rsa"
)

// Jwt Struct of Jwt keys
type Jwt struct {
	PrivateKey               *rsa.PrivateKey
	PublicKey                *rsa.PublicKey
	PublicKeyForInternalCall *rsa.PublicKey
}
