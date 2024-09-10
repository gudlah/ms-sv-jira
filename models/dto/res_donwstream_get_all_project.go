package dto

type ResDownstreamGetAllProject struct {
	ProjectId          int    `json:"projectId"`
	ProjectKey         string `json:"projectKey"`
	ProjectName        string `json:"projectName"`
	ProjectDisplayName string `json:"projectDisplayName"`
	ProjectTypeKey     string `json:"projectTypeKey"`
	// ProjectAvatarUrls  AvatarUrls `json:"projectavatarUrls"`
	ProjectAvatarUrls string `json:"projectavatarUrls"`
	BoardId           int    `json:"boardId"`
	BoardName         string `json:"boardName"`
	BoardType         string `json:"boardType"`
}
