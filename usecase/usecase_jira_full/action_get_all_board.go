package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) GetAllBoardAction(kosong interface{}, idRequest string) (httpCode int, res dto.Res, dataBoard dto.ResUpstreamGetAllBoard) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}
	resUpstream, err := usecase.ExternalRepository.GetAllBoardRepository()
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		logUpstream.IsSuccess = 1
		json.Unmarshal(resUpstream.Body(), &dataBoard)

		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataBoard)
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
