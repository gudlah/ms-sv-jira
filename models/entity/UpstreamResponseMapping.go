package entity

type UpstreamResponseMapping struct {
	Id                      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Endpoint                string `gorm:"type:varchar(255)" json:"endpoint"`
	ErrorCodeBrigate        string `gorm:"type:varchar(10)" json:"error_code_brigate"`
	ResponseCodeBrigate     string `gorm:"type:varchar(10)" json:"response_code_brigate"`
	ResponseMessageBrigate  string `gorm:"type:varchar(255)" json:"response_message_brigate"`
	KondisiInternalBribrain string `gorm:"type:varchar(255)" json:"kondisi_internal_bribrain"`
	ResponseCodeBribrain    string `gorm:"type:varchar(10)" json:"response_code_bribrain"`
	MessageBribrain         string `gorm:"type:varchar(255)" json:"message_bribrain"`
	IsSuccess               int    `gorm:"type:int(11)" json:"is_success"`
}
