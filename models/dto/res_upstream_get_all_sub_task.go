package dto

type ResUpstreamGetAllSubTask struct {
	Expand string                      `json:"expand"`
	ID     string                      `json:"id"`
	Self   string                      `json:"self"`
	Key    string                      `json:"key"`
	Fields FieldsUpstreamGetAllSubTask `json:"fields"`
}

type FieldsUpstreamGetAllSubTask struct {
	Subtasks []SubtaskUpstreamGetAllSubTask `json:"subtasks"`
}

type SubtaskUpstreamGetAllSubTask struct {
	ID     string                         `json:"id"`
	Key    string                         `json:"key"`
	Self   string                         `json:"self"`
	Fields SubFieldsUpstreamGetAllSubTask `json:"fields"`
}

type SubFieldsUpstreamGetAllSubTask struct {
	Summary   string                         `json:"summary"`
	Status    StatusUpstreamGetAllSubTask    `json:"status"`
	Priority  PriorityUpstreamGetAllSubTask  `json:"priority"`
	Issuetype IssuetypeUpstreamGetAllSubTask `json:"issuetype"`
}

type StatusUpstreamGetAllSubTask struct {
	Self           string                              `json:"self"`
	Description    string                              `json:"description"`
	IconURL        string                              `json:"iconUrl"`
	Name           string                              `json:"name"`
	ID             string                              `json:"id"`
	StatusCategory StatusCategoryUpstreamGetAllSubTask `json:"statusCategory"`
}

type StatusCategoryUpstreamGetAllSubTask struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type PriorityUpstreamGetAllSubTask struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type IssuetypeUpstreamGetAllSubTask struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	EntityID       string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}
