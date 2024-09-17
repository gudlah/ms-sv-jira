package usecase_jira

import (
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/usecase/usecase_log"
)

type JiraUsecaseImpl struct {
	ExternalRepository repository_external.ExternalRepository
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
}

func NewJiraUsecase(externalRepository repository_external.ExternalRepository, databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase) JiraUsecase {
	return &JiraUsecaseImpl{
		ExternalRepository: externalRepository,
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
	}
}
