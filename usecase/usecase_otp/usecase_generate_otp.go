package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"strings"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (usecase *OtpUsecaseImpl) GenerateOtpUsecase(kosong interface{}, idRequest, jenisTransaksi, phoneNumber string) (httpCode int, res dto.Res) {
	jenisTransaksi = strings.ToLower(jenisTransaksi)
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Sabrina",
		AccountName: phoneNumber,
		Period:      usecase.OtpExpiredSecond,
		Algorithm:   otp.AlgorithmSHA1,
	})

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		httpCode, res = usecase.GenerateOtpAction(kosong, key, idRequest, jenisTransaksi, phoneNumber)
	}

	return
}
