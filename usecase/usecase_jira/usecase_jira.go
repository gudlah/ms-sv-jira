package usecase_jira

import "ms-sv-jira/models/dto"

type JiraUsecase interface {
	GetAllUserUsecase(kosong interface{}, idRequest string) (int, dto.Res)
	GetAllProjectUsecase(kosong interface{}, idRequest string) (int, dto.Res)
	GetAllPriorityUsecase(kosong interface{}, idRequest string) (httpCode int, res dto.Res)
	GetAllSprintUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllSprint) (int, dto.Res)
	GetAllCardUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllCard) (int, dto.Res)
	GetAllSubTaskUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamGetAllSubTask) (int, dto.Res)
}
