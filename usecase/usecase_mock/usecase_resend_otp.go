package usecase_mock

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *MockUsecaseImpl) ResendOtpUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamResendOtp) (httpCode int, res dto.Res) {
	if bodyRequest.SessionId == "s4br1N1" {
		httpCode, res = helpers.ResReached("1010", "Request limit reached", kosong)
	} else if bodyRequest.SessionId == "s4br1N2" {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if bodyRequest.SessionId == "s4br1N3" {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		httpCode, res = usecase.OtpUsecase.GenerateOtpUsecase(kosong, idRequest, bodyRequest.JenisTransaksi, bodyRequest.PhoneNumber)
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
	}

	return
}
