package delivery_jira

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_jira"
	"ms-sv-jira/usecase/usecase_log"

	"github.com/go-playground/validator/v10"
)

type JiraDeliveryImpl struct {
	AuthDelivery delivery_auth.AuthDelivery
	LogUsecase   usecase_log.LogUsecase
	Validate     *validator.Validate
	JiraUsecase  usecase_jira.JiraUsecase
}

func NewJiraDelivery(authDelivery delivery_auth.AuthDelivery, jiraUsecase usecase_jira.JiraUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) JiraDelivery {
	return &JiraDeliveryImpl{
		AuthDelivery: authDelivery,
		LogUsecase:   logUsecase,
		Validate:     validate,
		JiraUsecase:  jiraUsecase,
	}
}
