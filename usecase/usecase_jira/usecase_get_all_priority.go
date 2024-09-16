package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraUsecaseImpl) GetAllPriorityUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}
	resUpstream, err := usecase.ExternalRepository.GetAllPriorityRepository()
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		resStruct := []dto.ResUpstreamGetAllPriority{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		logUpstream.IsSuccess = 1
		dataOutput := make([]dto.ResDownstreamGetAllPriority, len(resStruct))
		for indexPriority, priority := range resStruct {
			priorityIdInt, _ := strconv.Atoi(priority.Id)
			dataOutput[indexPriority] = dto.ResDownstreamGetAllPriority{
				PriorityId:          priorityIdInt,
				PriorityName:        priority.Name,
				PriorityDescription: priority.Description,
			}
		}
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
