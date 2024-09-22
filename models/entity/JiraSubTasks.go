package entity

import "time"

type JiraSubTasks struct {
	SubTaskID          string    `json:"sub_task_id" gorm:"column:sub_task_id;size:10;primaryKey"`
	CardKey            string    `json:"card_key,omitempty" gorm:"column:card_key;size:10"`
	SubTaskKey         string    `json:"sub_task_key,omitempty" gorm:"column:sub_task_key;size:10"`
	SubTaskTitle       string    `json:"sub_task_title,omitempty" gorm:"column:sub_task_title;size:100"`
	SubTaskDescription string    `json:"sub_task_description,omitempty" gorm:"column:sub_task_description;type:longtext"`
	StatusID           string    `json:"status_id,omitempty" gorm:"column:status_id;size:10"`
	PriorityID         int       `json:"priority_id,omitempty" gorm:"column:priority_id"`
	CreatorID          string    `json:"creator_id,omitempty" gorm:"column:creator_id;size:100"`
	ReporterID         string    `json:"reporter_id,omitempty" gorm:"column:reporter_id;size:100"`
	AssigneeID         string    `json:"assignee_id,omitempty" gorm:"column:assignee_id;size:100"`
	SubTaskCreated     time.Time `json:"sub_task_created,omitempty" gorm:"column:sub_task_created"`
	SubTaskUpdated     time.Time `json:"sub_task_updated,omitempty" gorm:"column:sub_task_updated"`
	SubTaskStarted     time.Time `json:"sub_task_started,omitempty" gorm:"column:sub_task_started"`
	SubTaskResolved    time.Time `json:"sub_task_resolved,omitempty" gorm:"column:sub_task_resolved"`
}
