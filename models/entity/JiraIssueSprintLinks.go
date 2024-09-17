package entity

import "time"

type JiraIssueSprintLinks struct {
	IssueSprintLinkID int        `gorm:"column:issue_sprint_link_id;primaryKey;autoIncrement"`
	IssueID           *string    `gorm:"column:issue_id;size:50"`
	SprintID          *string    `gorm:"column:sprint_id;size:50"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
}
