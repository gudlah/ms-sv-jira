package entity

type Symbols struct {
	Id     int    `json:"id" gorm:"type:integer;primary_key"`
	Symbol string `json:"symbol" gorm:"type:varchar(100)"`
}
