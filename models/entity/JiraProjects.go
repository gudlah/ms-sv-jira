package entity

type JiraProjects struct {
	ProjectID      string `json:"project_id" gorm:"column:project_id;size:50;primaryKey"`
	ProjectKey     string `json:"project_key,omitempty" gorm:"column:project_key;size:50"`
	ProjectName    string `json:"project_name,omitempty" gorm:"column:project_name;size:100"`
	ProjectTypeKey string `json:"project_type_key,omitempty" gorm:"column:project_type_key;size:50"`
}
