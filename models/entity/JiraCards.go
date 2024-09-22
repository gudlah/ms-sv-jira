package entity

import "time"

type JiraCards struct {
	CardID          string    `json:"card_id" gorm:"column:card_id;size:10;primaryKey"`
	ColumnID        string    `json:"column_id,omitempty" gorm:"column:column_id;size:10"`
	CardTitle       string    `json:"card_title,omitempty" gorm:"column:card_title;size:255"`
	CardKey         string    `json:"card_key,omitempty" gorm:"column:card_key;size:10"`
	CardDescription string    `json:"card_description,omitempty" gorm:"column:card_description;type:longtext"`
	CardTypeID      string    `json:"card_type_id,omitempty" gorm:"column:card_type_id;size:10"`
	PriorityID      int       `json:"priority_id,omitempty" gorm:"column:priority_id"`
	AssigneeID      string    `json:"assignee_id,omitempty" gorm:"column:assignee_id;size:100"`
	CreatorID       string    `json:"creator_id,omitempty" gorm:"column:creator_id;size:100"`
	ReporterID      string    `json:"reporter_id,omitempty" gorm:"column:reporter_id;size:100"`
	CardCreated     time.Time `json:"card_created,omitempty" gorm:"column:card_created"`
	CardUpdated     time.Time `json:"card_updated,omitempty" gorm:"column:card_updated"`
	CardStarted     time.Time `json:"card_started,omitempty" gorm:"column:card_started"`
	CardResolved    time.Time `json:"card_resolved,omitempty" gorm:"column:card_resolved"`
}
