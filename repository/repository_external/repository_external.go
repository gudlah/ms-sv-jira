package repository_external

import (
	"ms-sv-jira/models/dto"

	"github.com/go-resty/resty/v2"
)

type ExternalRepository interface {
	SmsRepository(body dto.ReqUpstreamSms) (*resty.Response, error)
}
