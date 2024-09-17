package dto

import (
	"ms-sv-jira/models/entity"
	"time"
)

type ResDownstreamGetAllProject struct {
	ProjectId          int                 `json:"projectId"`
	ProjectKey         string              `json:"projectKey"`
	ProjectName        string              `json:"projectName"`
	ProjectDescription string              `json:"projectDescription"`
	ProjectLeadName    string              `json:"projectLeadName"`
	ProjectType        string              `json:"projectType"`
	CreatedAt          time.Time           `json:"createdAt"`
	Created            time.Time           `json:"created"`
	Boards             []entity.JiraBoards `json:"boards"`
}
