package entity

import "time"

type JiraProjects struct {
	ProjectID      int        `gorm:"column:project_id;primaryKey;autoIncrement"`
	JiraProjectKey string     `gorm:"column:jira_project_key;size:50;not null"`
	Name           string     `gorm:"column:name;size:255;not null"`
	Description    *string    `gorm:"column:description;type:text"`
	Lead           *string    `gorm:"column:lead;size:255"`
	TypeProject    *string    `gorm:"column:type_project;size:255"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	Created        *time.Time `gorm:"column:created"`
}
