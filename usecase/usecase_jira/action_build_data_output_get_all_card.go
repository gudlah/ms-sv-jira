package usecase_jira

import (
	"ms-sv-jira/models/dto"
)

func (usecase *JiraUsecaseImpl) BuildDataOuputGetAllCardAction(kosong interface{}, idRequest string, sprintId int, dataColumn dto.ResUpstreamGetAllColumn, dataCard dto.ResUpstreamGetAllCard) (dataOutput []dto.ResDownstreamGetAllCard) {
	dataOutput = make([]dto.ResDownstreamGetAllCard, len(dataColumn.ColumnConfig.Columns))
	for indexColumn, column := range dataColumn.ColumnConfig.Columns {
		dataOutput[indexColumn] = dto.ResDownstreamGetAllCard{
			ColumnId:   column.Statuses[0].Id,
			SprintId:   sprintId,
			ColumnName: column.Name,
			Cards:      usecase.BuildDataCard(kosong, idRequest, column.Statuses[0].Id, column.Name, dataCard),
		}
	}
	return
}
