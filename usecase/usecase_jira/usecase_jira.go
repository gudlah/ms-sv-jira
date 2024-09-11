package usecase_jira

import "ms-sv-jira/models/dto"

type JiraUsecase interface {
	GetAllProjectUsecase(kosong interface{}, idRequest string) (int, dto.Res)
	GetAllSprintUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllSprint) (int, dto.Res)
	GetAllCardUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllCard) (int, dto.Res)
}
