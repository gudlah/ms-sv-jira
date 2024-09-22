package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) InsertJiraPriorityAction(kosong interface{}, priorities []dto.ResDownstreamGetAllPriority, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataPriorities := []entity.JiraPriorities{}
	var errGetPriority int
	for _, priority := range priorities {
		cekPriority, err := usecase.DatabaseRepository.GetJiraPrioritiesRepository(priority.PriorityId)
		dataInsertPriority := entity.JiraPriorities{
			PriorityID:          priority.PriorityId,
			PriorityName:        priority.PriorityName,
			PriorityDescription: priority.PriorityDescription,
		}
		if err != nil {
			errGetPriority += 1
		} else if cekPriority.PriorityID == 0 {
			dataPriorities = append(dataPriorities, dataInsertPriority)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraPrioritiesRepository(dataInsertPriority)
			if err != nil {
				errGetPriority += 1
			}
		}
	}

	if errGetPriority > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataPriorities) > 0 {
		err := usecase.DatabaseRepository.InsertJiraPrioritiesRepository(dataPriorities)
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
