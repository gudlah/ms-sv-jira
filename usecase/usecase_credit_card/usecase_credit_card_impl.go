package usecase_credit_card

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/usecase/usecase_log"
	"ms-sv-jira/usecase/usecase_otp"
)

type CreditCardUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
	ExternalRepository repository_external.ExternalRepository
	BrigateConfig      dto.BrigateConfig
	OtpUsecase         usecase_otp.OtpUsecase
}

func NewCreditCardUsecase(databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase, externalRepository repository_external.ExternalRepository, brigateConfig dto.BrigateConfig, otpUsecase usecase_otp.OtpUsecase) CreditCardUsecase {
	return &CreditCardUsecaseImpl{
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
		ExternalRepository: externalRepository,
		BrigateConfig:      brigateConfig,
		OtpUsecase:         otpUsecase,
	}
}
