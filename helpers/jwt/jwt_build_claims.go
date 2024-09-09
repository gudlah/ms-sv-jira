package helpers

import (
	"ms-sv-jira/helpers"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func BuildJWTClaims(ttl time.Duration, userId int) (claims jwt.MapClaims) {
	userIdString := strconv.Itoa(userId)
	encryptedUserId := helpers.Encrypt(userIdString)
	now := time.Now().UTC()
	claims = make(jwt.MapClaims)
	claims["dat"] = encryptedUserId
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	return
}
