package usecase_log

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *LogUsecaseImpl) InsertLogUpstreamUsecase(param dto.ParamLogUpstream) (httpCode int, response dto.Res) {
	err := usecase.DatabaseRepository.InsertLogUpstreamRepository(param.LogUpstream)
	if err != nil {
		httpCode, response = helpers.ResGeneralSystemError(param.Kosong)
	} else {
		httpCode, response = param.HttpCode, param.Response
	}
	return
}
