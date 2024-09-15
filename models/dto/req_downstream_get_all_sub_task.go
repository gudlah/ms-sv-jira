package dto

type ReqDownstreamGetAllSubTask struct {
	CardKey string `json:"cardKey" validate:"required"`
}
