package dto

import (
	"ms-sv-jira/models/entity"
)

type ParamLogUpstream struct {
	Response    Res
	LogUpstream entity.UpstreamServiceRequestLog
	Kosong      interface{}
	HttpCode    int
}
