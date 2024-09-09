package repository_external

import (
	"ms-sv-jira/config"
	"ms-sv-jira/models/dto"
	"time"

	"github.com/go-resty/resty/v2"
)

func (repository *ExternalRepositoryImpl) InquiryCustomerCreditCardRepository(body dto.ReqUpstreamInquiryCustomerCreditCard) (*resty.Response, error) {
	ginMode := config.Config.GinMode
	url := repository.BrigateConfig.EsbUrl + "/inquiryCreditCard"
	if ginMode != "lokal" {
		url = repository.BrigateConfig.EsbUrl + "/creditcardretrieve/1.0/mscreditcardretrieveinformation/getcustomercreditcard"
	}
	return repository.
		RestyConfig.
		SetTimeout(10*time.Second).
		R().
		SetBody(body).
		SetBasicAuth(repository.BrigateConfig.Username, repository.BrigateConfig.Password).
		Post(url)
}
