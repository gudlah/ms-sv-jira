package entity

import "time"

type JiraIssues struct {
	IssueID      int        `gorm:"column:issue_id;primaryKey;autoIncrement"`
	JiraIssueKey string     `gorm:"column:jira_issue_key;size:50;not null"`
	ProjectID    *string    `gorm:"column:project_id;size:50"`
	EpicKey      *string    `gorm:"column:epic_key;size:50"`
	Summary      string     `gorm:"column:summary;type:text;not null"`
	Description  *string    `gorm:"column:description;type:text"`
	Type         *string    `gorm:"column:type;size:50"`
	Status       *string    `gorm:"column:status;size:100"`
	AssigneeID   *string    `gorm:"column:assignee_id;size:50"`
	ReporterID   *string    `gorm:"column:reporter_id;size:50"`
	Priority     *string    `gorm:"column:priority;size:50"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
	Created      *time.Time `gorm:"column:created"`
	Updated      *time.Time `gorm:"column:updated"`
	ParentKey    *string    `gorm:"column:parent_key;size:50"`
}
