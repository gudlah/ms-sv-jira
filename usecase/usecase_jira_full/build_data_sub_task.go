package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"strconv"
)

func BuilDataSubTask(cardKey string, dataUpstream dto.ResUpstreamGetAllSubTask) (dataOutput []dto.SubTaskDownstreamGetAllFull) {
	dataOutput = make([]dto.SubTaskDownstreamGetAllFull, dataUpstream.Total)
	for index, subtask := range dataUpstream.Issues {
		field := subtask.Fields
		priorityIdInt, _ := strconv.Atoi(field.Priority.ID)
		dataOutput[index] = dto.SubTaskDownstreamGetAllFull{
			SubTaskId:          subtask.ID,
			CardKey:            cardKey,
			SubTaskKey:         subtask.Key,
			SubTaskTitle:       field.Summary,
			StatusId:           field.Status.ID,
			StatusName:         field.Status.Name,
			PriorityId:         priorityIdInt,
			PriorityName:       field.Priority.Name,
			Created:            helpers.KonversiTanggalIssueOutput(field.Created),
			Updated:            helpers.KonversiTanggalIssueOutput(field.Updated),
			Resolved:           helpers.KonversiTanggalIssueOutput(field.Resolutiondate),
			AssigneeId:         field.Assignee.AccountID,
			AssigneeName:       field.Assignee.DisplayName,
			CreatorId:          field.Creator.AccountID,
			CreatorName:        field.Creator.DisplayName,
			ReporterId:         field.Reporter.AccountID,
			ReporterName:       field.Reporter.DisplayName,
			SubTaskDescription: BuildDataSubTaskDescription(field.Description),
		}
	}
	return
}
