package helpers

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func BuildParamInsertLogUpstream(logUpstream entity.UpstreamServiceRequestLog, httpCode int, response dto.Res, kosong interface{}) dto.ParamLogUpstream {
	return dto.ParamLogUpstream{
		Response:    response,
		LogUpstream: logUpstream,
		Kosong:      kosong,
		HttpCode:    httpCode,
	}
}
