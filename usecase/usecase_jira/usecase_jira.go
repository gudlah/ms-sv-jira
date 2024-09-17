package usecase_jira

import "ms-sv-jira/models/dto"

type JiraUsecase interface {
	GetAllJiraProjectUsecase(kosong interface{}) (int, dto.Res)
	GetAllJiraUsersUsecase(kosong interface{}) (int, dto.Res)
}
