package dto

type Res struct {
	Success      bool        `json:"success"`
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}
