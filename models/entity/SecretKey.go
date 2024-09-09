package entity

type SecretKey struct {
	Id          string `json:"id" gorm:"type:varchar(100);primary_key"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20)"`
	SecretKey   string `json:"secret_key" gorm:"type:text"`
	Type        string `json:"type" gorm:"type:varchar(100)"`
	ExpiredAt   string `json:"expired_at" gorm:"type:varchar(100)"`
	CreatedAt   string `json:"created_at" gorm:"type:varchar(100)"`
	UpdatedAt   string `json:"updated_at" gorm:"type:varchar(100)"`
}
