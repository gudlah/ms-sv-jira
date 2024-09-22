package dto

type ResUpstreamIssueChangelog struct {
	Changelog Changelog `json:"changelog"`
}
type Changelog struct {
	Histories []History `json:"histories"`
}

type History struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Items   []Item `json:"items"`
}

type Item struct {
	Field      string `json:"field"`
	FieldType  string `json:"fieldtype"`
	FieldID    string `json:"fieldId"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}
