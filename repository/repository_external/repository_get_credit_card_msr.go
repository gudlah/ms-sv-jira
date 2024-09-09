package repository_external

import (
	"ms-sv-jira/config"
	"ms-sv-jira/models/dto"
	"time"

	"github.com/go-resty/resty/v2"
)

func (repository *ExternalRepositoryImpl) GetCreditCardMsrRepository(body dto.ReqUpstreamGetCreditCardMsr) (*resty.Response, error) {
	ginMode := config.Config.GinMode
	url := repository.BrigateConfig.EsbUrl + "/getCreditCard"
	if ginMode != "lokal" {
		url = repository.BrigateConfig.EsbUrl + "/creditcardretrieve/1.0/mscreditcardretrieveinformation/getcreditcard"
	}
	return repository.
		RestyConfig.
		SetTimeout(10*time.Second).
		R().
		SetBody(body).
		SetBasicAuth(repository.BrigateConfig.Username, repository.BrigateConfig.Password).
		Post(url)
}
