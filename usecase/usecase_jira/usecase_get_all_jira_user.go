package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraUsecaseImpl) GetAllJiraUsersUsecase(kosong interface{}) (httpCode int, res dto.Res) {
	dataUsers, err := usecase.DatabaseRepository.GetAllJiraUsersRepository()
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataUsers)
	}
	return
}
