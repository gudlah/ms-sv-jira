package usecase_jwt

import (
	"ms-sv-jira/config"
	jwtHelpers "ms-sv-jira/helpers/jwt"
	"ms-sv-jira/models/dto"
	"time"
)

func (usecase *JWTUsecaseImpl) GenerateJWTUsecase(userId int) (res dto.ResTokenJWT, err error) {
	expiredTime := time.Duration(config.Config.ExpiredJWT)
	durasi := time.Minute * expiredTime
	token, err := jwtHelpers.CreateTokenJWT(usecase.SecretKey.PrivateKey, durasi, userId)
	if err == nil {
		res = jwtHelpers.BuildOutputJWT(token, durasi)
	}
	return
}
