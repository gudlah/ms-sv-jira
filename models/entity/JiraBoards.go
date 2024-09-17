package entity

import "time"

type JiraBoards struct {
	BoardID     int        `json:"board_id" gorm:"column:board_id;primaryKey;autoIncrement"`
	JiraBoardID string     `json:"jira_board_id" gorm:"column:jira_board_id;size:50;not null"`
	ProjectID   *string    `json:"project_id,omitempty" gorm:"column:project_id;size:50"`
	Name        string     `json:"name" gorm:"column:name;size:255;not null"`
	Type        *string    `json:"type,omitempty" gorm:"column:type;size:50"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	Created     time.Time  `json:"created" gorm:"column:created;autoCreateTime"`
}
