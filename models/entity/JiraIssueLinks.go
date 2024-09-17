package entity

import "time"

type JiraIssueLinks struct {
	LinkID       int       `gorm:"column:link_id;primaryKey;autoIncrement"`
	IssueID      *string   `gorm:"column:issue_id;size:50"`
	LinkIssueKey *string   `gorm:"column:link_issue_key;size:50"`
	URL          *string   `gorm:"column:url;size:500"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	Title        *string   `gorm:"column:title;size:50"`
}
