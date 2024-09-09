package usecase_credit_card

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *CreditCardUsecaseImpl) CheckLimitUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckLimit) (httpCode int, res dto.Res) {
	jumlahSessionId, err := usecase.LogUsecase.GetJumlahSessionUsecase(bodyRequest.SessionId, "transaction")
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if jumlahSessionId >= 3 {
		httpCode, res = helpers.ResSuccess(false, "1009", "OTP limit reached", kosong)
	} else {
		paramVerifOtp := dto.ParamVerifOtp{
			PhoneNumber:    bodyRequest.PhoneNumber,
			OtpCode:        bodyRequest.OtpCode,
			JenisTransaksi: "limit",
		}
		httpCode, res = usecase.OtpUsecase.VerifOtpUsecase(kosong, paramVerifOtp)
		if res.ResponseCode == "0000" {
			httpCode, res = usecase.CheckLimitAction(idRequest, kosong, bodyRequest)
		}
	}

	return

}
