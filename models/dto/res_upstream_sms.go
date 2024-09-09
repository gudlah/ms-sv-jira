package dto

type ResUpstreamSms struct {
	ErrorCode       string `json:"errorCode"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReffNumber      string `json:"reffNumber"`
}
