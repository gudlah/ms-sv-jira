package dto

type ReqDownstreamGetAllCard struct {
	BoardId  int `json:"boardId" validate:"required,numeric"`
	SprintId int `json:"sprintId" validate:"required,numeric"`
}
