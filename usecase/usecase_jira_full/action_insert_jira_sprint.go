package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) InsertJiraSprintAction(kosong interface{}, sprints []dto.SprintDownstreamGetAllFull, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataSprints := []entity.JiraSprints{}
	var errGetSprint int
	for _, sprint := range sprints {
		cekSprint, err := usecase.DatabaseRepository.GetJiraSprintRepository(sprint.SprintId)
		dataInsertSprint := entity.JiraSprints{
			SprintID:          sprint.SprintId,
			BoardID:           sprint.BoardId,
			SprintState:       sprint.State,
			SprintName:        sprint.Name,
			SprintStartDate:   helpers.KonversiTanggalDb(sprint.StartDate),
			SprintEndDate:     helpers.KonversiTanggalDb(sprint.EndDate),
			SprintCreatedDate: helpers.KonversiTanggalDb(sprint.CreatedDate),
			SprintGoal:        sprint.Goal,
		}
		if err != nil {
			errGetSprint += 1
		} else if cekSprint.SprintID == 0 {
			dataSprints = append(dataSprints, dataInsertSprint)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraSprintRepository(dataInsertSprint)
			if err != nil {
				errGetSprint += 1
			}
		}
	}

	if errGetSprint > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataSprints) > 0 {
		err := usecase.DatabaseRepository.InsertJiraSprintRepository(dataSprints)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = httpCodeAsal, resAsal
		}
	} else {
		httpCode, res = httpCodeAsal, resAsal
	}

	return
}
