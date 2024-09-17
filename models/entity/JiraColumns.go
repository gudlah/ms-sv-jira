package entity

type JiraColumns struct {
	ColumnID   string `json:"column_id" gorm:"column:column_id;size:10;primaryKey"`
	SprintID   int    `json:"sprint_id,omitempty" gorm:"column:sprint_id;size:11"`
	ColumnName string `json:"column_name,omitempty" gorm:"column:column_name;size:100"`
}
