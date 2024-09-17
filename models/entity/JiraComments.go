package entity

import "time"

type JiraComments struct {
	CommentID int        `gorm:"column:comment_id;primaryKey;autoIncrement"`
	IssueID   *string    `gorm:"column:issue_id;size:50"`
	AuthorID  *string    `gorm:"column:author_id;size:50"`
	Body      *string    `gorm:"column:body;type:text"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	Created   time.Time  `gorm:"column:created;autoCreateTime"`
}
