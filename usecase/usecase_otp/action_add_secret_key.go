package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *OtpUsecaseImpl) AddSecretKeyAction(kosong interface{}, phoneNumber, secretKey, jenisTransaksi, expiredAt string) (httpCode int, res dto.Res) {
	sekarang := helpers.Now()
	secretData := entity.SecretKey{
		Id:          helpers.GenerateUUID(),
		PhoneNumber: phoneNumber,
		SecretKey:   secretKey,
		Type:        jenisTransaksi,
		CreatedAt:   sekarang,
		ExpiredAt:   expiredAt,
		UpdatedAt:   sekarang,
	}
	err := usecase.DatabaseRepository.InsertSecretKeyRepository(secretData)
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
	}

	return
}
