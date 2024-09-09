package delivery_mock

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_log"
	"ms-sv-jira/usecase/usecase_mock"

	"github.com/go-playground/validator/v10"
)

type MockDeliveryImpl struct {
	AuthDelivery delivery_auth.AuthDelivery
	MockUsecase  usecase_mock.MockUsecase
	LogUsecase   usecase_log.LogUsecase
	Validate     *validator.Validate
}

func NewMockDelivery(authDelivery delivery_auth.AuthDelivery, mockUsecase usecase_mock.MockUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) MockDelivery {
	return &MockDeliveryImpl{
		AuthDelivery: authDelivery,
		MockUsecase:  mockUsecase,
		LogUsecase:   logUsecase,
		Validate:     validate,
	}
}
