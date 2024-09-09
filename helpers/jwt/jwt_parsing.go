package helpers

import (
	"crypto/rsa"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func ParsingJwt(signedToken string, key *rsa.PublicKey) (token *jwt.Token, err error) {
	token, err = jwt.Parse(signedToken, func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodRSA)
		if !ok {
			err = fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			return nil, err
		} else {
			return key, nil
		}
	})
	return
}
