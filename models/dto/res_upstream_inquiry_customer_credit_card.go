package dto

type ResUpstreamInquiryCustomerCreditCard struct {
	Response ResponseUpstreamInquiryCustomerCreditCard `json:"response"`
}

type ResponseUpstreamInquiryCustomerCreditCard struct {
	ErrorCode       string `json:"errorCode"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	HomePhone       string `json:"homePhone"`
	OfficePhone     string `json:"officePhone"`
	EmployerName    string `json:"employerName"`
	CustomerStatus  string `json:"customerStatus"`
}
