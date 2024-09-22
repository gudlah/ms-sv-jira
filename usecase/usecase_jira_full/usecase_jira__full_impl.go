package usecase_jira_full

import (
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/usecase/usecase_log"
)

type JiraFullUsecaseImpl struct {
	ExternalRepository repository_external.ExternalRepository
	DatabaseRepository repository_database.DatabaseRepository
	LogUsecase         usecase_log.LogUsecase
}

func NewJiraFullUsecase(externalRepository repository_external.ExternalRepository, databaseRepository repository_database.DatabaseRepository, logUsecase usecase_log.LogUsecase) JiraFullUsecase {
	return &JiraFullUsecaseImpl{
		ExternalRepository: externalRepository,
		DatabaseRepository: databaseRepository,
		LogUsecase:         logUsecase,
	}
}
