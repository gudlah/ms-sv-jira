package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/pquerna/otp/totp"
)

func (usecase *OtpUsecaseImpl) VerifOtpUsecase(kosong interface{}, param dto.ParamVerifOtp) (httpCode int, res dto.Res) {
	secretKeyData, err := usecase.DatabaseRepository.GetSecretKeyRepository(param.JenisTransaksi, param.PhoneNumber)

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if secretKeyData.Id == "" {
		httpCode, res = helpers.ResSuccess(false, "1003", "OTP Invalid", kosong)
	} else {
		isValid := totp.Validate(param.OtpCode, secretKeyData.SecretKey)
		if !isValid {
			httpCode, res = helpers.ResSuccess(false, "1003", "OTP Invalid", kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
		}
	}

	return
}
