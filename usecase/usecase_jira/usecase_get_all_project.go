package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) GetAllProjectUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}
	resUpstream, err := usecase.ExternalRepository.GetAllProjetRepository()
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		logUpstream.IsSuccess = 1
		resStruct := []dto.ResUpstreamGetAllProject{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		dataOutput := make([]dto.ResDownstreamGetAllProject, len(resStruct))

		for i, dataUpstream := range resStruct {
			dataOutput[i] = dto.ResDownstreamGetAllProject{
				ProjectId:         dataUpstream.Id,
				ProjectKey:        dataUpstream.Key,
				ProjectName:       dataUpstream.Name,
				ProjectTypeKey:    dataUpstream.ProjectTypeKey,
				ProjectAvatarUrls: dataUpstream.AvatarUrls,
			}
		}

		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
