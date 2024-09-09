package repository_external

import (
	"ms-sv-jira/models/dto"

	"github.com/go-resty/resty/v2"
)

type ExternalRepository interface {
	InquiryCustomerCreditCardRepository(body dto.ReqUpstreamInquiryCustomerCreditCard) (*resty.Response, error)
	GetCreditCardMsrRepository(body dto.ReqUpstreamGetCreditCardMsr) (*resty.Response, error)
	SmsRepository(body dto.ReqUpstreamSms) (*resty.Response, error)
}
