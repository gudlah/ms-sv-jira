package entity

import "time"

type JiraAttachments struct {
	AttachmentID       string    `json:"attachment_id" gorm:"column:attachment_id;size:10;primaryKey"`
	CardID             string    `json:"card_id,omitempty" gorm:"column:card_id;size:10"`
	AttachmentFileName string    `json:"attachment_file_name,omitempty" gorm:"column:attachment_file_name;size:100"`
	AttachmentURL      string    `json:"attachment_url,omitempty" gorm:"column:attachment_url;type:text"`
	AuthorID           string    `json:"author_id,omitempty" gorm:"column:author_id;size:100"`
	AttachmentCreated  time.Time `json:"attachment_created,omitempty" gorm:"column:attachment_created"`
}
