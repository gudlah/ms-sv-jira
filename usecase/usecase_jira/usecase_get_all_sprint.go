package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraUsecaseImpl) GetAllSprintUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllSprint) (httpCode int, res dto.Res, dataOutput []dto.ResDownstreamGetAllSprint) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	boardId := strconv.Itoa(bodyRequest.BoardId)
	resUpstream, err := usecase.ExternalRepository.GetAllSprintRepository(boardId)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		resStruct := dto.ResUpstreamGetAllSprint{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		if resStruct.Total == 0 {
			logUpstream.IsSuccess = 0
			httpCode, res = helpers.ResSuccess(true, "1003", "Data not found", kosong)
		} else {
			logUpstream.IsSuccess = 1
			dataOutput = BuildDataSprint(resStruct)
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
			httpCode, res = usecase.InsertJiraSprintAction(kosong, dataOutput, httpCode, res)
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
