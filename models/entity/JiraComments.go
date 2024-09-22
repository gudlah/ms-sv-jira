package entity

import "time"

type JiraComments struct {
	CommentID      string    `json:"comment_id" gorm:"column:comment_id;size:10;primaryKey"`
	CardID         string    `json:"card_id,omitempty" gorm:"column:card_id;size:10"`
	AuthorID       string    `json:"author_id,omitempty" gorm:"column:author_id;size:100"`
	CommentBody    string    `json:"comment_body,omitempty" gorm:"column:comment_body;type:longtext"`
	CommentCreated time.Time `json:"comment_created,omitempty" gorm:"column:comment_created"`
	CommentUpdated time.Time `json:"comment_updated,omitempty" gorm:"column:comment_updated"`
}
