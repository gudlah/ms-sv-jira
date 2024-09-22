package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraFullUsecaseImpl) BuildDataSprint(kosong interface{}, idRequest string, boardId int, sprint dto.ResUpstreamGetAllSprint) (httpCode int, res dto.Res, dataOutput []dto.SprintDownstreamGetAllFull) {
	dataOutput = make([]dto.SprintDownstreamGetAllFull, len(sprint.Values))
	for index, sprint := range sprint.Values {
		httpCodeColumn, resColumn, dataColumn := usecase.GetAllColumnAction(kosong, idRequest, boardId)
		if resColumn.ResponseCode != "0000" {
			httpCode, res = httpCodeColumn, resColumn
			return
		}

		httpCodeCard, resCard, dataCard := usecase.GetAllCardAction(kosong, idRequest, sprint.Id)
		if resCard.ResponseCode != "0000" {
			httpCode, res = httpCodeCard, resCard
			return
		}
		httpCodeDataColumn, resDataColumn, columns := usecase.BuildDataColumn(kosong, idRequest, sprint.Id, dataColumn, dataCard)
		if resDataColumn.ResponseCode != "0000" {
			httpCode, res = httpCodeDataColumn, resDataColumn
			return
		}
		dataOutput[index] = dto.SprintDownstreamGetAllFull{
			SprintId:    sprint.Id,
			BoardId:     sprint.OriginBoardId,
			State:       sprint.State,
			Name:        sprint.Name,
			StartDate:   helpers.KonversiTanggalOutput(sprint.StartDate),
			EndDate:     helpers.KonversiTanggalOutput(sprint.EndDate),
			CreatedDate: helpers.KonversiTanggalOutput(sprint.CreatedDate),
			Goal:        sprint.Goal,
			Columns:     columns,
		}
	}

	return
}
