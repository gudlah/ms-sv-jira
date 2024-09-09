package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *AuthUsecaseImpl) ValidateBlockUsecase(ipClient string, httpCode int, res dto.Res, kosong interface{}) (httpCodeHasil int, resHasil dto.Res) {
	ipBlocked, err := usecase.DatabaseRepository.GetBlockedIpRepository(ipClient)
	if err != nil {
		httpCodeHasil, resHasil = helpers.ResGeneralSystemError(kosong)
	} else if ipBlocked.Id != "" {
		httpCodeHasil, resHasil = helpers.ResIpBlocked(kosong)
	} else {
		httpCodeHasil, resHasil = usecase.AksiBelumDiblockUsecase(ipClient, httpCode, res, kosong)
	}
	return
}
