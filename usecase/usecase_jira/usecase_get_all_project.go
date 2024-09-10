package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraUsecaseImpl) GetAllProjectUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res) {
	httpCode, res, dataProject := usecase.GetAllProjectAction(kosong, idRequest)
	if res.ResponseCode != "0000" {
		return
	}

	httpCode, res, dataBoard := usecase.GetAllBoardAction(kosong, idRequest)
	if res.ResponseCode != "0000" {
		return
	}

	dataOutput := usecase.BuildDataOutputGetAllProjectAction(dataProject, dataBoard)
	httpCode, res = helpers.ResSuccess(true, "0000", "Successfuly", dataOutput)
	return
}
