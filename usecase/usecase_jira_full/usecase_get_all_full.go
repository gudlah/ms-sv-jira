package usecase_jira_full

import (
	"ms-sv-jira/models/dto"
)

func (usecase *JiraFullUsecaseImpl) GetAllFullUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res) {
	httpCode, res = usecase.BuildDataProject(kosong, idRequest)
	return
}
