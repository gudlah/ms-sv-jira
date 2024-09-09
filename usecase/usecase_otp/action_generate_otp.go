package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (usecase *OtpUsecaseImpl) GenerateOtpAction(kosong interface{}, key *otp.Key, idRequest, jenisTransaksi, phoneNumber string) (httpCode int, res dto.Res) {
	secretKey := key.Secret()
	durasi := time.Duration(key.Period())
	expiredTime := time.Now().Add(time.Second * durasi)
	otp, err := totp.GenerateCode(secretKey, expiredTime)

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		httpCode, res = usecase.CheckAvailabilitySecretKeyAction(kosong, jenisTransaksi, phoneNumber, secretKey, expiredTime)
	}

	if res.ResponseCode == "0000" {
		httpCode, res = usecase.SmsAction(kosong, idRequest, phoneNumber, otp, jenisTransaksi)
	}

	return
}
