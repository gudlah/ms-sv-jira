package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraFullUsecaseImpl) GetAllColumnAction(kosong interface{}, idRequest string, boardId int) (httpCode int, res dto.Res, resStruct dto.ResUpstreamGetAllColumn) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}

	boardIdString := strconv.Itoa(boardId)
	resUpstream, err := usecase.ExternalRepository.GetAllColumnRepository(boardIdString)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		json.Unmarshal(resUpstream.Body(), &resStruct)

		jumlahColumn := len(resStruct.ColumnConfig.Columns)
		if resStruct.Id == 0 {
			logUpstream.IsSuccess = 0
			httpCode, res = helpers.ResSuccess(false, "1003", "Data not found", kosong)
		} else if jumlahColumn == 0 {
			logUpstream.IsSuccess = 1
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
