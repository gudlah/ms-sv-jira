package usecase_jwt

import (
	"ms-sv-jira/models/dto"
)

type JWTUsecase interface {
	GenerateJWTUsecase(userId int) (res dto.ResTokenJWT, err error)
	ValidateTokenUsecase(signedToken string) (string, error)
}
