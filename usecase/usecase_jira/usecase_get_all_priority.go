package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
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
		logUpstream.IsSuccess = 1
		logUpstream.ResponsePayload = resUpstream.String()
		resStruct := []dto.ResUpstreamGetAllPriority{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		dataOutput := BuildDataPriority(resStruct)
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		httpCode, res = usecase.InsertJiraPriorityAction(kosong, dataOutput, httpCode, res)
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
