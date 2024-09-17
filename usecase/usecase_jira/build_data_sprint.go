package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func BuildDataSprint(sprint dto.ResUpstreamGetAllSprint) (dataOutput []dto.ResDownstreamGetAllSprint) {
	dataOutput = make([]dto.ResDownstreamGetAllSprint, len(sprint.Values))
	for index, sprint := range sprint.Values {
		dataOutput[index] = dto.ResDownstreamGetAllSprint{
			SprintId:    sprint.Id,
			BoardId:     sprint.OriginBoardId,
			State:       sprint.State,
			Name:        sprint.Name,
			StartDate:   helpers.KonversiTanggalOutput(sprint.StartDate),
			EndDate:     helpers.KonversiTanggalOutput(sprint.EndDate),
			CreatedDate: helpers.KonversiTanggalOutput(sprint.CreatedDate),
			Goal:        sprint.Goal,
		}
	}

	return
}
