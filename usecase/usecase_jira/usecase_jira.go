package usecase_jira

import "ms-sv-jira/models/dto"

type JiraUsecase interface {
	GetAllProjectUsecase(kosong interface{}, idRequest string) (int, dto.Res)
}
