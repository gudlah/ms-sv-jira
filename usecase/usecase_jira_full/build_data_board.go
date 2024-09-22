package usecase_jira_full

import (
	"ms-sv-jira/models/dto"
	"strconv"
)

func (usecase *JiraFullUsecaseImpl) BuildDataBoard(kosong interface{}, idRequest string, projectId string, dataBoard dto.ResUpstreamGetAllBoard) (hasilBoards []dto.BoardDownstreamGetAllFull) {
	for _, board := range dataBoard.Values {
		idProjectInt, _ := strconv.Atoi(projectId)
		if board.Location.ProjectId == idProjectInt {
			_, _, sprints := usecase.GetDataSprintAction(kosong, idRequest, board.Id)
			if len(sprints) < 1 {
				sprints = make([]dto.SprintDownstreamGetAllFull, 0)
			}
			hasilBoards = append(hasilBoards, dto.BoardDownstreamGetAllFull{
				BoardId:   board.Id,
				ProjectId: projectId,
				BoardName: board.Name,
				BoardType: board.Type,
				Sprints:   sprints,
			})
		}
	}
	return
}
