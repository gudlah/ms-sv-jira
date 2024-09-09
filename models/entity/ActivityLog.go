package entity

type ActivityLog struct {
	Id           string      `json:"id" gorm:"type:varchar(100);primary_key"`
	IdRequest    string      `json:"id_request" gorm:"type:varchar(100)"`
	IdUser       interface{} `json:"id_user" gorm:"type:varchar(100)"`
	ClientIp     string      `json:"client_ip" gorm:"type:varchar(100)"`
	SessionId    string      `json:"session_id" gorm:"type:varchar(255)"`
	Activity     string      `json:"activity" gorm:"type:varchar(100)"`
	ResponseCode string      `json:"response_code" gorm:"type:varchar(100)"`
	BodyRequest  string      `json:"body_request" gorm:"type:text"`
	BodyResponse string      `json:"body_response" gorm:"type:text;default:{}"`
	CreatedAt    string      `json:"created_at" gorm:"type:datetime"`
}
