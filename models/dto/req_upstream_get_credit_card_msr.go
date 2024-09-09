package dto

type ReqUpstreamGetCreditCardMsr struct {
	Request RequestUpstreamGetCreditCardMsr `json:"request"`
}

type RequestUpstreamGetCreditCardMsr struct {
	CardNumber string `json:"cardNumber"`
	ChannelId  string `json:"channelId"`
	ExternalId string `json:"externalId"`
	ServiceId  string `json:"serviceId"`
}
