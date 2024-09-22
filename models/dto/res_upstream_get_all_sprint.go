package dto

type ResUpstreamGetAllSprint struct {
	MaxResults int                         `json:"maxResults"`
	StartAt    int                         `json:"startAt"`
	Total      int                         `json:"total"`
	IsLast     bool                        `json:"isLast"`
	Values     []ValueUpstreamGetAllSprint `json:"values"`
}

type ValueUpstreamGetAllSprint struct {
	Id            int    `json:"id"`
	Self          string `json:"self"`
	State         string `json:"state"`
	Name          string `json:"name"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	CreatedDate   string `json:"createdDate"`
	OriginBoardId int    `json:"originBoardId"`
	Goal          string `json:"goal"`
}
