package entity

type UpstreamServiceRequestLog struct {
	Id                string `json:"id" gorm:"type:varchar(100);primary_key"`
	IdRequest         string `json:"id_request" gorm:"type:varchar(100)"`
	RequestPayload    string `json:"request_payload" gorm:"type:text"`
	ResponsePayload   string `json:"response_payload" gorm:"type:text"`
	IsSuccess         int    `json:"is_success" gorm:"type:int"`
	RequestTimestamp  string `json:"request_timestamp" gorm:"type:datetime"`
	ResponseTimestamp string `json:"response_timestamp" gorm:"type:datetime"`
	Url               string `json:"url" gorm:"type:varchar(100)"`
}
