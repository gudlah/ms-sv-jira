package usecase_auth

import (
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/usecase/usecase_log"
)

type AuthUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
}

func NewAuthUsecase(databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase) AuthUsecase {
	return &AuthUsecaseImpl{
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
	}
}
