package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraFullUsecaseImpl) BuildDataColumn(kosong interface{}, idRequest string, sprintId int, dataColumn dto.ResUpstreamGetAllColumn, dataCard dto.ResUpstreamGetAllCard) (httpCode int, res dto.Res, dataOutput []dto.ColumnDownstreamGetAllFull) {
	dataOutput = make([]dto.ColumnDownstreamGetAllFull, len(dataColumn.ColumnConfig.Columns))
	for indexColumn, column := range dataColumn.ColumnConfig.Columns {
		httpCodeCard, resCard, cards := usecase.BuildDataCard(kosong, idRequest, column.Statuses[0].Id, column.Name, dataCard)
		if resCard.ResponseCode != "0000" {
			httpCode, res = httpCodeCard, resCard
			return
		} else {
			dataOutput[indexColumn] = dto.ColumnDownstreamGetAllFull{
				ColumnId:   column.Statuses[0].Id,
				SprintId:   sprintId,
				ColumnName: column.Name,
				Cards:      cards,
			}
		}
	}
	if httpCode == 0 {
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		httpCode, res = usecase.InsertJiraColumnAction(kosong, dataOutput, httpCode, res)
	}
	return
}
