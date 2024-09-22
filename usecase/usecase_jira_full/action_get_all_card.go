package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraFullUsecaseImpl) GetAllCardAction(kosong interface{}, idRequest string, sprintId int) (httpCode int, res dto.Res, resStruct dto.ResUpstreamGetAllCard) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	sprintIdString := strconv.Itoa(sprintId)
	resUpstream, err := usecase.ExternalRepository.GetAllCardRepository(sprintIdString)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		json.Unmarshal(resUpstream.Body(), &resStruct)

		if resStruct.MaxResults == 0 {
			logUpstream.IsSuccess = 0
			httpCode, res = helpers.ResSuccess(false, "1003", "Data not found", kosong)
		} else {
			logUpstream.IsSuccess = 1
			httpCode, res = helpers.ResSuccess(false, "0000", "Successfully", kosong)
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
