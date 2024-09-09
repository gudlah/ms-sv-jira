package repository_external

import (
	"ms-sv-jira/config"
	"ms-sv-jira/models/dto"
	"time"

	"github.com/go-resty/resty/v2"
)

func (repository *ExternalRepositoryImpl) SmsRepository(body dto.ReqUpstreamSms) (*resty.Response, error) {
	ginMode := config.Config.GinMode
	url := repository.BrigateConfig.EsbUrl + "/smsNotification"
	if ginMode != "lokal" {
		url = repository.BrigateConfig.EsbUrl + "/BRINotification/2.0/SmsNotification"
	}
	return repository.
		RestyConfig.
		SetTimeout(10*time.Second).
		R().
		SetBody(body).
		SetHeader("appname", repository.BrigateConfig.SmsAppName).
		SetBasicAuth(repository.BrigateConfig.Username, repository.BrigateConfig.Password).
		Post(url)
}
