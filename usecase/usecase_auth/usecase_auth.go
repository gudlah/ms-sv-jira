package usecase_auth

import (
	"ms-sv-jira/models/dto"
)

type AuthUsecase interface {
	LoginUsecase(kosong interface{}, loginRequest dto.ReqLogin) (int, dto.Res, string)
	ValidateSpecialCharUsecase(kosong interface{}, input string, statusCode int, response dto.Res) (int, dto.Res)
	UnblockIpUsecase(req dto.ReqUnblockIp, kosong interface{}) (int, dto.Res)
	ValidateBlockUsecase(ipClient string, httpCode int, res dto.Res, kosong interface{}) (int, dto.Res)
}
