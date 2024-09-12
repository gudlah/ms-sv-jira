package dto

type ReqDownstreamGetAllSubTask struct {
	CardId string `json:"cardId" validate:"required,numeric"`
}
