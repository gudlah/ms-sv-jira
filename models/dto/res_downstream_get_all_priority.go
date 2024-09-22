package dto

type ResDownstreamGetAllPriority struct {
	PriorityId          int    `json:"priorityId"`
	PriorityName        string `json:"priorityName"`
	PriorityDescription string `json:"priorityDescription"`
}
