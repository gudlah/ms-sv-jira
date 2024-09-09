package delivery_jira

import (
	"ms-sv-jira/usecase/usecase_auth"
	"ms-sv-jira/usecase/usecase_jira"
	"ms-sv-jira/usecase/usecase_log"

	"github.com/go-playground/validator/v10"
)

type JiraDeliveryImpl struct {
	AuthUsecase usecase_auth.AuthUsecase
	LogUsecase  usecase_log.LogUsecase
	Validate    *validator.Validate
	JiraUsecase usecase_jira.JiraUsecase
}

func NewJiraDelivery(authUsecase usecase_auth.AuthUsecase, jiraUsecase usecase_jira.JiraUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) JiraDelivery {
	return &JiraDeliveryImpl{
		AuthUsecase: authUsecase,
		LogUsecase:  logUsecase,
		Validate:    validate,
		JiraUsecase: jiraUsecase,
	}
}
