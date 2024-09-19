package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) InsertJiraColumnAction(kosong interface{}, columns []dto.ColumnDownstreamGetAllFull, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataColumns := []entity.JiraColumns{}
	var errGetColumn int
	for _, column := range columns {
		cekColumn, err := usecase.DatabaseRepository.GetJiraColumnRepository(column.ColumnId)
		dataInsertColumn := entity.JiraColumns{
			ColumnID:   column.ColumnId,
			ColumnName: column.ColumnName,
			SprintID:   column.SprintId,
		}
		if err != nil {
			errGetColumn += 1
		} else if cekColumn.ColumnID == "" {
			dataColumns = append(dataColumns, dataInsertColumn)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraColumnRepository(dataInsertColumn)
			if err != nil {
				errGetColumn += 1
			}
		}
	}

	if errGetColumn > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataColumns) > 0 {
		err := usecase.DatabaseRepository.InsertJiraColumnRepository(dataColumns)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = httpCodeAsal, resAsal
		}
	} else {
		httpCode, res = httpCodeAsal, resAsal
	}

	var errInsertCard int
	for _, column := range columns {
		_, resCard := usecase.InsertJiraCardAction(kosong, column.Cards, httpCode, res)
		if resCard.ResponseCode != "0000" {
			errInsertCard += 1
		}
	}

	return
}
