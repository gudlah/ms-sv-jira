package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraUsecaseImpl) GetAllCardUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllCard) (httpCode int, res dto.Res) {
	httpCode, res, dataColumn := usecase.GetAllColumnAction(kosong, idRequest, bodyRequest.BoardId)
	if res.ResponseCode != "0000" {
		return
	}

	httpCode, res, dataCard := usecase.GetAllCardAction(kosong, idRequest, bodyRequest.SprintId)
	if res.ResponseCode != "0000" {
		return
	}

	dataOutput := usecase.BuildDataOuputGetAllCardAction(dataColumn, dataCard)
	httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	return
}
