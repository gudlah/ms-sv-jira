package entity

import "time"

type JiraSprints struct {
	SprintID     int        `gorm:"column:sprint_id;primaryKey;autoIncrement"`
	JiraSprintID string     `gorm:"column:jira_sprint_id;size:50;not null"`
	BoardID      *string    `gorm:"column:board_id;size:50"`
	Name         *string    `gorm:"column:name;size:255"`
	Goal         *string    `gorm:"column:goal;type:text"`
	State        *string    `gorm:"column:state;size:50"`
	StartDate    *time.Time `gorm:"column:start_date"`
	EndDate      *time.Time `gorm:"column:end_date"`
	CompleteDate *time.Time `gorm:"column:complete_date"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	Created      time.Time  `gorm:"column:created;autoCreateTime"`
}
