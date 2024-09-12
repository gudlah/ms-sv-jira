package dto

type ResDownstreamGetAllSubTask struct {
	SubTaskId    string `json:"subTaskId"`
	SubTaskKey   string `json:"subTaskKey"`
	SubTaskTitle string `json:"subTaskTitle"`
	StatusId     string `json:"statusId"`
	StatusName   string `json:"statusName"`
	PriorityId   string `json:"priorityId"`
	PriorityName string `json:"priorityName"`
}
