package dto

type ResDownstreamGetAllProject struct {
	ProjectId      string                         `json:"projectId"`
	ProjectKey     string                         `json:"projectKey"`
	ProjectName    string                         `json:"projectName"`
	ProjectTypeKey string                         `json:"projectTypeKey"`
	Boards         []BoardDownstreamGetAllProject `json:"boards"`
}

type BoardDownstreamGetAllProject struct {
	BoardId   int    `json:"boardId"`
	ProjectId string `json:"projectId"`
	BoardName string `json:"boardName"`
	BoardType string `json:"boardType"`
}
