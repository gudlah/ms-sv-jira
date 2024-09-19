package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *JiraFullUsecaseImpl) BuildDataProject(kosong interface{}, idRequest string) (httpCode int, res dto.Res) {
	httpCode, res, dataProject := usecase.GetAllProjectAction(kosong, idRequest)
	if res.ResponseCode != "0000" {
		return
	}

	httpCode, res, dataBoard := usecase.GetAllBoardAction(kosong, idRequest)
	if res.ResponseCode != "0000" {
		return
	}

	dataOutput := make([]dto.ResDownstreamGetAllFull, len(dataProject))

	for indexProject, project := range dataProject {
		boards := usecase.BuildDataBoard(kosong, idRequest, project.Id, dataBoard)
		if len(boards) < 1 {
			boards = make([]dto.BoardDownstreamGetAllFull, 0)
		}
		httpCode, res = usecase.InsertJiraBoardAction(kosong, boards, httpCode, res)
		dataOutput[indexProject] = dto.ResDownstreamGetAllFull{
			ProjectId:      project.Id,
			ProjectKey:     project.Key,
			ProjectName:    project.Name,
			ProjectTypeKey: project.ProjectTypeKey,
			Boards:         boards,
		}
	}
	httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	httpCode, res = usecase.InsertJiraProjectAction(kosong, dataOutput, httpCode, res)

	return
}
