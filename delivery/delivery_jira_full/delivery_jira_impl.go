package delivery_jira_full

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_jira_full"
	"ms-sv-jira/usecase/usecase_log"

	"github.com/go-playground/validator/v10"
)

type JiraFullDeliveryImpl struct {
	AuthDelivery    delivery_auth.AuthDelivery
	LogUsecase      usecase_log.LogUsecase
	Validate        *validator.Validate
	JiraFullUsecase usecase_jira_full.JiraFullUsecase
}

func NewJiraFullDelivery(authDelivery delivery_auth.AuthDelivery, jiraFullUsecase usecase_jira_full.JiraFullUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) JiraFullDelivery {
	return &JiraFullDeliveryImpl{
		AuthDelivery:    authDelivery,
		LogUsecase:      logUsecase,
		Validate:        validate,
		JiraFullUsecase: jiraFullUsecase,
	}
}
