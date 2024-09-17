package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) GetAllJiraProjectUsecase(kosong interface{}) (httpCode int, res dto.Res) {
	dataProjects, err := usecase.DatabaseRepository.GetAllJiraProjectRepository()
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		dataOutput := make([]dto.ResDownstreamGetAllProject, len(dataProjects))
		var errBoard int
		for indexProject, project := range dataProjects {
			dataOutput[indexProject] = dto.ResDownstreamGetAllProject{
				ProjectId:          project.ProjectID,
				ProjectKey:         project.JiraProjectKey,
				ProjectName:        project.Name,
				ProjectDescription: *project.Description,
				ProjectLeadName:    *project.Lead,
				ProjectType:        *project.TypeProject,
				CreatedAt:          *project.CreatedAt,
				Created:            *project.Created,
			}
			dataBoards, err := usecase.DatabaseRepository.GetAllJiraBoardsRepository(project.JiraProjectKey)
			if err != nil {
				errBoard += 1
			} else if len(dataBoards) == 0 {
				dataOutput[indexProject].Boards = make([]entity.JiraBoards, 0)
			} else {
				dataOutput[indexProject].Boards = dataBoards
			}
		}
		if errBoard > 0 {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		}
	}
	return
}
