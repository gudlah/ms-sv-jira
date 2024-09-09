package usecase_auth

import (
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/usecase/usecase_jwt"
	"ms-sv-jira/usecase/usecase_log"
)

type AuthUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
	JWTUsecase         usecase_jwt.JWTUsecase
}

func NewAuthUsecase(databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase, jwtUsecase usecase_jwt.JWTUsecase) AuthUsecase {
	return &AuthUsecaseImpl{
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
		JWTUsecase:         jwtUsecase,
	}
}
