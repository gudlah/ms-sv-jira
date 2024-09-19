package usecase_jira

import (
	"ms-sv-jira/models/dto"
	"strconv"
)

func BuildDataOutputGetAllProject(dataProject []dto.ResUpstreamGetAllProject, dataBoard dto.ResUpstreamGetAllBoard) (dataOutput []dto.ResDownstreamGetAllProject) {
	dataOutput = make([]dto.ResDownstreamGetAllProject, len(dataProject))

	for indexProject, project := range dataProject {
		hasilBoards := []dto.BoardDownstreamGetAllProject{}
		for _, board := range dataBoard.Values {
			idProjectInt, _ := strconv.Atoi(project.Id)
			if board.Location.ProjectId == idProjectInt {
				hasilBoards = append(hasilBoards, dto.BoardDownstreamGetAllProject{
					BoardId:   board.Id,
					ProjectId: project.Id,
					BoardName: board.Name,
					BoardType: board.Type,
				})
			}
		}

		dataOutput[indexProject] = dto.ResDownstreamGetAllProject{
			ProjectId:      project.Id,
			ProjectKey:     project.Key,
			ProjectName:    project.Name,
			ProjectTypeKey: project.ProjectTypeKey,
			Boards:         hasilBoards,
		}
	}

	return
}
