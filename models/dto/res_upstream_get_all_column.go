package dto

type ResUpstreamGetAllColumn struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Self         string       `json:"self"`
	Location     Location     `json:"location"`
	Filter       Filter       `json:"filter"`
	ColumnConfig ColumnConfig `json:"columnConfig"`
	Estimation   Estimation   `json:"estimation"`
	Ranking      Ranking      `json:"ranking"`
}

type Location struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	ID   string `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
}

type Filter struct {
	ID   string `json:"id"`
	Self string `json:"self"`
}

type ColumnConfig struct {
	Columns        []Column `json:"columns"`
	ConstraintType string   `json:"constraintType"`
}

type Statuses struct {
	Id   string `json:"id"`
	Self string `json:"self"`
}

type Column struct {
	Name     string     `json:"name"`
	Statuses []Statuses `json:"statuses"`
}

type Estimation struct {
	Type  string `json:"type"`
	Field Field  `json:"field"`
}

type Field struct {
	FieldID     string `json:"fieldId"`
	DisplayName string `json:"displayName"`
}

type Ranking struct {
	RankCustomFieldID int `json:"rankCustomFieldId"`
}
