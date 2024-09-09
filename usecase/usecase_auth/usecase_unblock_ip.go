package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *AuthUsecaseImpl) UnblockIpUsecase(req dto.ReqUnblockIp, kosong interface{}) (httpCode int, res dto.Res) {
	blockedIp, err := usecase.DatabaseRepository.GetBlockedIpRepository(req.IpAddress)
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if blockedIp.Id == "" {
		httpCode, res = helpers.ResSuccess(false, "1003", "IP address not found", kosong)
	} else {
		err = usecase.DatabaseRepository.SoftDeleteBlockedIpRepository(blockedIp.Id)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
		}
	}
	return
}
