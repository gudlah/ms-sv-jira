package usecase_jira

import (
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/usecase/usecase_log"
)

type JiraUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
}

func NewJiraUsecase(databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase) JiraUsecase {
	return &JiraUsecaseImpl{
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
	}
}
