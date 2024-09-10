package dto

type ReqDownstreamGetAllSprint struct {
	BoardId int `json:"boardId" validate:"required,numeric"`
}
