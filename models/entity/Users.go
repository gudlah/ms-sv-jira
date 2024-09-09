package entity

type Users struct {
	Id       int    `json:"id" gorm:"type:integer;primary_key"`
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(100)"`
	IsActive int    `json:"is_active" gorm:"type:integer(1)"`
}
