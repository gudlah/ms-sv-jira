package usecase_jwt

import (
	"fmt"
	"ms-sv-jira/helpers"
	jwtHelpers "ms-sv-jira/helpers/jwt"

	jwt "github.com/dgrijalva/jwt-go"
)

func (usecase *JWTUsecaseImpl) ValidateTokenUsecase(signedToken string) (userId string, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(usecase.SecretKey.PublicKey)
	if err != nil {
		return "0", err
	} else {
		tok, err := jwtHelpers.ParsingJwt(signedToken, key)
		if err != nil {
			return "0", err
		} else {
			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok || !tok.Valid {
				err = fmt.Errorf("validate: invalid")
				return "0", err
			} else {
				datString := helpers.InterfaceToString(claims["dat"])
				userId, err := helpers.Decrypt(string(datString))
				if err != nil {
					return "0", err
				} else {
					return userId, nil
				}
			}
		}
	}
}
