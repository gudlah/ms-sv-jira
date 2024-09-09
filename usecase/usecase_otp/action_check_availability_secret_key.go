package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"time"
)

func (usecase *OtpUsecaseImpl) CheckAvailabilitySecretKeyAction(kosong interface{}, jenisTransaksi, phoneNumber, secretKey string, expiredTime time.Time) (httpCode int, res dto.Res) {
	dataSecretKey, err := usecase.DatabaseRepository.GetSecretKeyRepository(jenisTransaksi, phoneNumber)
	expiredAtString := expiredTime.Format("2006-01-02 15:04:05 MST")

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if dataSecretKey.Id == "" {
		httpCode, res = usecase.AddSecretKeyAction(kosong, phoneNumber, secretKey, jenisTransaksi, expiredAtString)
	} else {
		httpCode, res = usecase.UpdateSecretKeyAction(kosong, dataSecretKey, secretKey, expiredAtString)
	}

	return
}
