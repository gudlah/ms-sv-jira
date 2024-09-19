package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) InsertJiraProjectAction(kosong interface{}, projects []dto.ResDownstreamGetAllFull, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataProjects := []entity.JiraProjects{}
	var errProject int
	for _, project := range projects {
		cekProject, err := usecase.DatabaseRepository.GetJiraProjectRepository(project.ProjectId)
		dataInsertProject := entity.JiraProjects{
			ProjectID:      project.ProjectId,
			ProjectKey:     project.ProjectKey,
			ProjectName:    project.ProjectName,
			ProjectTypeKey: project.ProjectTypeKey,
		}
		if err != nil {
			errProject += 1
		} else if cekProject.ProjectID == "" {
			dataProjects = append(dataProjects, dataInsertProject)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraProjectRepository(dataInsertProject)
			if err != nil {
				errProject += 1
			}
		}
	}

	if errProject > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataProjects) > 0 {
		err := usecase.DatabaseRepository.InsertJiraProjectRepository(dataProjects)
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
