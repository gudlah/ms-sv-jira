package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *OtpUsecaseImpl) UpdateSecretKeyAction(kosong interface{}, dataSecretKey entity.SecretKey, secretKey, expiredAt string) (httpCode int, res dto.Res) {
	secretData := entity.SecretKey{
		Id:          dataSecretKey.Id,
		PhoneNumber: dataSecretKey.PhoneNumber,
		SecretKey:   secretKey,
		Type:        dataSecretKey.Type,
		CreatedAt:   dataSecretKey.CreatedAt,
		ExpiredAt:   expiredAt,
		UpdatedAt:   helpers.Now(),
	}
	err := usecase.DatabaseRepository.UpdateSecretKeyRepository(secretData)
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
	}

	return
}
