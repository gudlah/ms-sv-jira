package entity

type JiraPriorities struct {
	PriorityID          int    `json:"priority_id" gorm:"column:priority_id;primaryKey;autoIncrement"`
	PriorityName        string `json:"priority_name,omitempty" gorm:"column:priority_name;size:100"`
	PriorityDescription string `json:"priority_description,omitempty" gorm:"column:priority_description;type:text"`
}
