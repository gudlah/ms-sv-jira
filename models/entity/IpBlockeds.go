package entity

type IpBlockeds struct {
	Id        string `json:"id" gorm:"type:varchar(255);primary_key"`
	ClientIp  string `json:"client_ip" gorm:"type:varchar(255)"`
	CreatedAt string `json:"created_at" gorm:"type:varchar(100)"`
	DeletedAt string `json:"deleted_at" gorm:"type:varchar(100)"`
}
