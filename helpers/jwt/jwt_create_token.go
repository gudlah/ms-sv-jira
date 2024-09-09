package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateTokenJWT(privateKey []byte, ttl time.Duration, userId int) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	} else {
		claims := BuildJWTClaims(ttl, userId)
		token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
		if err != nil {
			return "", err
		} else {
			return token, nil
		}
	}
}
