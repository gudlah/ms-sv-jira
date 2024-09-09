package usecase_log

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *LogUsecaseImpl) InsertLogActivityUsecase(param dto.ActivityLogParam, sessionId string) (int, dto.Res) {
	activityLogBaru := helpers.BuildActivityLog(param, sessionId)
	err := usecase.DatabaseRepository.InsertLogActivityRepository(activityLogBaru)
	if err != nil {
		param.HttpCode, param.Response = helpers.ResGeneralSystemError(param.Kosong)
	}
	return param.HttpCode, param.Response
}
