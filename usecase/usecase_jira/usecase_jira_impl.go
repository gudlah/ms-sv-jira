package usecase_jira

import (
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/usecase/usecase_log"
)

type JiraUsecaseImpl struct {
	ExternalRepository repository_external.ExternalRepository
	LogUsecase         usecase_log.LogUsecase
}

func NewJiraUsecase(externalRepository repository_external.ExternalRepository, logUsecase usecase_log.LogUsecase) JiraUsecase {
	return &JiraUsecaseImpl{
		ExternalRepository: externalRepository,
		LogUsecase:         logUsecase,
	}
}
