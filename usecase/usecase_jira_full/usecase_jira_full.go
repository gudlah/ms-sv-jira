package usecase_jira_full

import "ms-sv-jira/models/dto"

type JiraFullUsecase interface {
	GetAllFullUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res)
}
