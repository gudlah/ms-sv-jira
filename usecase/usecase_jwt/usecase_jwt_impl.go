package usecase_jwt

import (
	"ms-sv-jira/models/dto"

	"github.com/go-playground/validator/v10"
)

type JWTUsecaseImpl struct {
	Validate  *validator.Validate
	SecretKey dto.SecretKey
}

func NewJWTUsecase(validate *validator.Validate, secretKey dto.SecretKey) JWTUsecase {
	return &JWTUsecaseImpl{
		Validate:  validate,
		SecretKey: secretKey,
	}
}
