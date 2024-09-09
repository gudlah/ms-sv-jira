package usecase_otp

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/usecase/usecase_log"
)

type OtpUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
	ExternalRepository repository_external.ExternalRepository
	BrigateConfig      dto.BrigateConfig
	OtpExpiredSecond   uint
}

func NewOtpUsecase(databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase, externalRepository repository_external.ExternalRepository, brigateConfig dto.BrigateConfig, otpExpiredSecond uint) OtpUsecase {
	return &OtpUsecaseImpl{
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
		ExternalRepository: externalRepository,
		BrigateConfig:      brigateConfig,
		OtpExpiredSecond:   otpExpiredSecond,
	}
}
