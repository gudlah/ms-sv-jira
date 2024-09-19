package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraFullUsecaseImpl) GetDataSprintAction(kosong interface{}, idRequest string, boardId int) (httpCode int, res dto.Res, dataOutput []dto.SprintDownstreamGetAllFull) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	boardIdString := strconv.Itoa(boardId)
	resUpstream, err := usecase.ExternalRepository.GetAllSprintRepository(boardIdString)
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
			httpCode, res, dataOutput = usecase.BuildDataSprint(kosong, idRequest, boardId, resStruct)
			if res.ResponseCode == "0000" {
				httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
			}
			httpCode, res = usecase.InsertJiraSprintAction(kosong, dataOutput, httpCode, res)
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
