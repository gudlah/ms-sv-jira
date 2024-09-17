package entity

import "time"

type Attachment struct {
	AttachmentID     int        `gorm:"column:attachment_id;primaryKey;autoIncrement"`
	IssueID          *string    `gorm:"column:issue_id;size:50"`
	FileName         *string    `gorm:"column:file_name;size:255"`
	MimeType         *string    `gorm:"column:mime_type;size:100"`
	FileSize         *int       `gorm:"column:file_size"`
	JiraAttachmentID *string    `gorm:"column:jira_attachment_id;size:50"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	Created          *time.Time `gorm:"column:created"`
}
