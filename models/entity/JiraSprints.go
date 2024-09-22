package entity

import "time"

type JiraSprints struct {
	SprintID          int       `json:"sprint_id" gorm:"column:sprint_id;primaryKey;autoIncrement"`
	BoardID           int       `json:"board_id,omitempty" gorm:"column:board_id"`
	SprintState       string    `json:"sprint_state,omitempty" gorm:"column:sprint_state;size:100"`
	SprintName        string    `json:"sprint_name,omitempty" gorm:"column:sprint_name;size:100"`
	SprintStartDate   time.Time `json:"sprint_start_date,omitempty" gorm:"column:sprint_start_date"`
	SprintEndDate     time.Time `json:"sprint_end_date,omitempty" gorm:"column:sprint_end_date"`
	SprintCreatedDate time.Time `json:"sprint_created_date,omitempty" gorm:"column:sprint_created_date"`
	SprintGoal        string    `json:"sprint_goal,omitempty" gorm:"column:sprint_goal;size:100"`
}
