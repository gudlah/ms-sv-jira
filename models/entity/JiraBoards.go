package entity

type JiraBoards struct {
	BoardID   int    `json:"board_id" gorm:"column:board_id;primaryKey;autoIncrement"`
	ProjectID string `json:"project_id,omitempty" gorm:"column:project_id;size:100"`
	BoardName string `json:"board_name,omitempty" gorm:"column:board_name;size:100"`
	BoardType string `json:"board_type,omitempty" gorm:"column:board_type;size:10"`
}
