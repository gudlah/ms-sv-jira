package dto

type ResUpstreamGetAllSubTask struct {
	Expand     string                       `json:"expand"`
	StartAt    int                          `json:"startAt"`
	MaxResults int                          `json:"maxResults"`
	Total      int                          `json:"total"`
	Issues     []IssueUpstreamGetAllSubTask `json:"issues"`
}

type IssueUpstreamGetAllSubTask struct {
	ID     string                      `json:"id"`
	Expand string                      `json:"expand"`
	Key    string                      `json:"key"`
	Fields FieldsUpstreamGetAllSubTask `json:"fields"`
}

type FieldsUpstreamGetAllSubTask struct {
	StatusCategoryChangeDate string                           `json:"statuscategorychangedate"`
	IssueType                TypeUpstreamGetAllSubTask        `json:"issuetype"`
	Created                  string                           `json:"created"`
	Priority                 PriorityUpstreamGetAllSubTask    `json:"priority"`
	Assignee                 UserUpstreamGetAllSubTask        `json:"assignee"`
	Updated                  string                           `json:"updated"`
	Resolutiondate           string                           `json:"resolutiondate"`
	Status                   StatusUpstreamGetAllSubTask      `json:"status"`
	Summary                  string                           `json:"summary"`
	Creator                  UserUpstreamGetAllSubTask        `json:"creator"`
	Reporter                 UserUpstreamGetAllSubTask        `json:"reporter"`
	Description              DescriptionUpstreamGetAllSubTask `json:"description"`
}

type TypeUpstreamGetAllSubTask struct {
	ID string `json:"id"`
}

type PriorityUpstreamGetAllSubTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUpstreamGetAllSubTask struct {
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

type StatusUpstreamGetAllSubTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DescriptionUpstreamGetAllSubTask struct {
	Type    string                         `json:"type"`
	Content []ContentUpstreamGetAllSubTask `json:"content"`
}

type ContentUpstreamGetAllSubTask struct {
	Type    string                              `json:"type"`
	Text    string                              `json:"text"`
	Content []ValueContentUpstreamGetAllSubTask `json:"content"`
}

type ValueContentUpstreamGetAllSubTask struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
