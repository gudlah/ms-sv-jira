package dto

type LocationUpstreamGetAllBoard struct {
	ProjectId      int    `json:"projectId"`
	DisplayName    string `json:"displayName"`
	ProjectName    string `json:"projectName"`
	ProjectKey     string `json:"projectKey"`
	ProjectTypeKey string `json:"projectTypeKey"`
	AvatarURI      string `json:"avatarURI"`
	Name           string `json:"name"`
}

type ValueUpstreamGetAllBoard struct {
	Id       int                         `json:"id"`
	Self     string                      `json:"self"`
	Name     string                      `json:"name"`
	Type     string                      `json:"type"`
	Location LocationUpstreamGetAllBoard `json:"location"`
}

type ResUpstreamGetAllBoard struct {
	MaxResults int                        `json:"maxResults"`
	StartAt    int                        `json:"startAt"`
	Total      int                        `json:"total"`
	IsLast     bool                       `json:"isLast"`
	Values     []ValueUpstreamGetAllBoard `json:"values"`
}
