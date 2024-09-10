package usecase_jira

import (
	"ms-sv-jira/models/dto"
	"strconv"
)

func (usecase *JiraUsecaseImpl) BuildDataOutputGetAllProjectAction(dataProject []dto.ResUpstreamGetAllProject, dataBoard dto.ResUpstreamGetAllBoard) (dataOutput []dto.ResDownstreamGetAllProject) {
	dataOutput = make([]dto.ResDownstreamGetAllProject, len(dataProject))

	for indexProject, project := range dataProject {
		idProjectInt, _ := strconv.Atoi(project.Id)
		dataOutput[indexProject] = dto.ResDownstreamGetAllProject{
			ProjectId:      idProjectInt,
			ProjectKey:     project.Key,
			ProjectName:    project.Name,
			ProjectTypeKey: project.ProjectTypeKey,
		}

		for _, board := range dataBoard.Values {
			if board.Location.ProjectId == idProjectInt {
				dataOutput[indexProject].ProjectDisplayName = board.Location.DisplayName
				dataOutput[indexProject].BoardId = board.Id
				dataOutput[indexProject].BoardName = board.Name
				dataOutput[indexProject].BoardType = board.Type
				dataOutput[indexProject].ProjectAvatarUrls = board.Location.AvatarURI
			}
		}
	}

	return
}
