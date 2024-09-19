package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) GetDataSubTaskAction(kosong interface{}, idRequest, cardKey string) (httpCode int, res dto.Res, dataOutput []dto.SubTaskDownstreamGetAllFull) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	resUpstream, err := usecase.ExternalRepository.GetAllSubtaskRepository(cardKey)
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

		// if resStruct.MaxResults == 0 {
		// 	logUpstream.IsSuccess = 0
		// 	httpCode, res = helpers.ResSuccess(true, "1003", "Data not found", kosong)
		// } else {
		logUpstream.IsSuccess = 1
		dataOutput = BuilDataSubTask(cardKey, resStruct)
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		// httpCode, res = usecase.InsertJiraSubTaskAction(kosong, dataOutput, httpCode, res)
		// }
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)
	return
}
