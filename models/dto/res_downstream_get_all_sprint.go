package dto

type ResDownstreamGetAllSprint struct {
	SprintId    int    `json:"sprintId"`
	BoardId     int    `json:"boardId"`
	State       string `json:"state"`
	Name        string `json:"name"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	CreatedDate string `json:"createdDate"`
	Goal        string `json:"goal"`
}
