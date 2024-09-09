package usecase_log

import (
	"ms-sv-jira/models/dto"
)

type LogUsecase interface {
	InsertLogActivityUsecase(param dto.ActivityLogParam) (int, dto.Res)
	InsertLogUpstreamUsecase(param dto.ParamLogUpstream) (httpCode int, response dto.Res)
	GetLastClientHitUsecase(ipClient string) (int, error)
}
