package entity

import "time"

type JiraUsers struct {
	UserID      int       `json:"user_id" gorm:"column:user_id;primaryKey;autoIncrement"`
	JiraUserID  string    `json:"jira_user_id" gorm:"column:jira_user_id;size:50;not null"`
	DisplayName *string   `json:"display_name,omitempty" gorm:"column:display_name;size:255"`
	Email       *string   `json:"email,omitempty" gorm:"column:email;size:255"`
	Active      bool      `json:"active" gorm:"column:active;default:1"`
	Created     time.Time `json:"created" gorm:"column:created;autoCreateTime"`
}
