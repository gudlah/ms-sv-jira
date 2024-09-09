package delivery_credit_card

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_credit_card"
	"ms-sv-jira/usecase/usecase_log"

	"github.com/go-playground/validator/v10"
)

type CreditCardDeliveryImpl struct {
	AuthDelivery      delivery_auth.AuthDelivery
	CreditCardUsecase usecase_credit_card.CreditCardUsecase
	LogUsecase        usecase_log.LogUsecase
	Validate          *validator.Validate
}

func NewCreditCardDelivery(authDelivery delivery_auth.AuthDelivery, creditCardUsecase usecase_credit_card.CreditCardUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) CreditCardDelivery {
	return &CreditCardDeliveryImpl{
		AuthDelivery:      authDelivery,
		CreditCardUsecase: creditCardUsecase,
		LogUsecase:        logUsecase,
		Validate:          validate,
	}
}
