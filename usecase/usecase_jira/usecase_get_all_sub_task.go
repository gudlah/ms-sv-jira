package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraUsecaseImpl) GetAllSubTaskUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllSubTask) (httpCode int, res dto.Res) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	resUpstream, err := usecase.ExternalRepository.GetAllSubtaskRepository(bodyRequest.CardKey)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		resStruct := dto.ResUpstreamGetAllSubTask{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		if resStruct.MaxResults == 0 {
			logUpstream.IsSuccess = 0
			httpCode, res = helpers.ResSuccess(true, "1003", "Data not found", kosong)
		} else {
			logUpstream.IsSuccess = 1
			dataOutput := builDataSubTask(resStruct)
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}

func builDataSubTask(dataUpstream dto.ResUpstreamGetAllSubTask) (dataOutput []dto.ResDownstreamGetAllSubTask) {
	dataOutput = make([]dto.ResDownstreamGetAllSubTask, dataUpstream.Total)
	for index, subtask := range dataUpstream.Issues {
		field := subtask.Fields
		priorityIdInt, _ := strconv.Atoi(field.Priority.ID)
		dataOutput[index] = dto.ResDownstreamGetAllSubTask{
			SubTaskId:          subtask.ID,
			SubTaskKey:         subtask.Key,
			SubTaskTitle:       field.Summary,
			StatusId:           field.Status.ID,
			StatusName:         field.Status.Name,
			PriorityId:         priorityIdInt,
			PriorityName:       field.Priority.Name,
			Created:            field.Created,
			Updated:            field.Updated,
			Resolved:           field.Resolutiondate,
			AssigneeId:         field.Assignee.AccountID,
			AssigneeName:       field.Assignee.DisplayName,
			CreatorId:          field.Creator.AccountID,
			CreatorName:        field.Creator.DisplayName,
			ReporterId:         field.Reporter.AccountID,
			ReporterName:       field.Reporter.DisplayName,
			SubTaskDescription: buildDescription(field.Description),
		}
	}
	return
}

func buildDescription(dataDescription dto.DescriptionUpstreamGetAllSubTask) (outputDescription string) {
	if dataDescription.Type != "" {
		for _, descriptionContent := range dataDescription.Content {
			if descriptionContent.Type == "paragraph" {
				for _, valueDescriptionContent := range descriptionContent.Content {
					if valueDescriptionContent.Type == "text" {
						outputDescription = valueDescriptionContent.Text
					}
				}
			}
		}
	}
	return
}
