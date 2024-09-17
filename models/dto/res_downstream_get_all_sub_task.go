package dto

type ResDownstreamGetAllSubTask struct {
	SubTaskId          string `json:"subTaskId"`
	CardKey            string `json:"cardKey"`
	SubTaskKey         string `json:"subTaskKey"`
	SubTaskTitle       string `json:"subTaskTitle"`
	SubTaskDescription string `json:"subTaskDescription"`
	StatusId           string `json:"statusId"`
	StatusName         string `json:"statusName"`
	PriorityId         int    `json:"priorityId"`
	PriorityName       string `json:"priorityName"`
	Created            string `json:"created"`
	Updated            string `json:"updated"`
	Resolved           string `json:"resolved"`
	CreatorId          string `json:"creatorId"`
	CreatorName        string `json:"creatorName"`
	ReporterId         string `json:"reporterId"`
	ReporterName       string `json:"reporterName"`
	AssigneeId         string `json:"assigneeId"`
	AssigneeName       string `json:"assigneeName"`
}
