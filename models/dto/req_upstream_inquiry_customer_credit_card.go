package dto

type ReqUpstreamInquiryCustomerCreditCard struct {
	Request RequestUpstreamInquiryCustomerCreditCard `json:"request"`
}

type RequestUpstreamInquiryCustomerCreditCard struct {
	CardNumber string `json:"cardNumber"`
	ChannelId  string `json:"channelId"`
	ExternalId string `json:"externalId"`
}
