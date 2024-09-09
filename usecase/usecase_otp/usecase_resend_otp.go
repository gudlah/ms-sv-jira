package usecase_otp

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *OtpUsecaseImpl) ResendOtpUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamResendOtp) (httpCode int, res dto.Res) {

	jumlahSessionId, err := usecase.LogUsecase.GetJumlahSessionUsecase(bodyRequest.SessionId, "otp")
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if jumlahSessionId >= 2 {
		httpCode, res = helpers.ResSuccess(false, "1010", "Request limit reached", kosong)
	} else {
		httpCode, res = usecase.GenerateOtpUsecase(kosong, idRequest, bodyRequest.JenisTransaksi, bodyRequest.PhoneNumber)
	}

	return
}
