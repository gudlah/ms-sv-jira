package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) InsertJiraSubTaskAction(kosong interface{}, subTasks []dto.ResDownstreamGetAllSubTask, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataSubTasks := []entity.JiraSubTasks{}
	var errGetSubTask int
	for _, subTask := range subTasks {
		cekSubTask, err := usecase.DatabaseRepository.GetJiraSubTaskRepository(subTask.SubTaskId)
		dataInsertSubTask := entity.JiraSubTasks{
			SubTaskID:          subTask.SubTaskId,
			CardKey:            subTask.CardKey,
			SubTaskKey:         subTask.SubTaskKey,
			SubTaskTitle:       subTask.SubTaskTitle,
			SubTaskDescription: subTask.SubTaskDescription,
			StatusID:           subTask.StatusId,
			PriorityID:         subTask.PriorityId,
			CreatorID:          subTask.CreatorId,
			ReporterID:         subTask.ReporterId,
			AssigneeID:         subTask.AssigneeId,
			SubTaskCreated:     helpers.KonversiTanggalDb(subTask.Created),
			SubTaskUpdated:     helpers.KonversiTanggalDb(subTask.Updated),
			SubTaskResolved:    helpers.KonversiTanggalDb(subTask.Resolved),
			SubTaskStarted:     helpers.KonversiTanggalDb(subTask.Started),
		}
		if err != nil {
			errGetSubTask += 1
		} else if cekSubTask.SubTaskID == "" {
			dataSubTasks = append(dataSubTasks, dataInsertSubTask)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraSubTaskRepository(dataInsertSubTask)
			if err != nil {
				errGetSubTask += 1
			}
		}
	}

	if errGetSubTask > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataSubTasks) > 0 {
		err := usecase.DatabaseRepository.InsertJiraSubTaskRepository(dataSubTasks)
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
