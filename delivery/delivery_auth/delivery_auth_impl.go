package delivery_auth

import (
	"ms-sv-jira/usecase/usecase_auth"
	"ms-sv-jira/usecase/usecase_log"

	"github.com/go-playground/validator/v10"
)

type AuthDeliveryImpl struct {
	AuthUsecase usecase_auth.AuthUsecase
	LogUsecase  usecase_log.LogUsecase
	Validate    *validator.Validate
}

func NewAuthDelivery(authUsecase usecase_auth.AuthUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) AuthDelivery {
	return &AuthDeliveryImpl{
		AuthUsecase: authUsecase,
		LogUsecase:  logUsecase,
		Validate:    validate,
	}
}
