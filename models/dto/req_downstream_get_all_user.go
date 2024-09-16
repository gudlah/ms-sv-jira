package dto

type ReqDownstreamGetAlUser struct {
	StartAt   int `json:"startAt" validate:"numeric"`
	MaxResult int `json:"maxResult" validate:"numeric"`
}
