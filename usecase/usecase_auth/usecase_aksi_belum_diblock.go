package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *AuthUsecaseImpl) AksiBelumDiblockUsecase(ipClient string, httpCode int, res dto.Res, kosong interface{}) (httpCodeHasil int, resHasil dto.Res) {
	jumlahHitGagal, err := usecase.LogUsecase.GetLastClientHitUsecase(ipClient)
	if err != nil {
		httpCodeHasil, resHasil = helpers.ResGeneralSystemError(kosong)
	} else if jumlahHitGagal >= 2 && httpCode != 200 {
		ipBlock := entity.IpBlockeds{
			Id:        helpers.GenerateUUID(),
			ClientIp:  ipClient,
			CreatedAt: helpers.Now(),
		}
		err = usecase.DatabaseRepository.InsertIpBlockRepository(ipBlock)
		if err != nil {
			httpCodeHasil, resHasil = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCodeHasil, resHasil = helpers.ResInvalidCredential(kosong)
		}
	} else {
		httpCodeHasil, resHasil = httpCode, res
	}
	return
}
