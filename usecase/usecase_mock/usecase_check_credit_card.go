package usecase_mock

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *MockUsecaseImpl) CheckCreditCardUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamCheckCreditCard) (httpCode int, res dto.Res) {
	if bodyRequest.CardNumber == "4359650100001201" {
		httpCode, res = helpers.ResSuccess(false, "1003", "Data not found", kosong)
	} else if bodyRequest.CardNumber == "4359650100001202" {
		httpCode, res = helpers.ResSuccess(false, "1004", "Card status closed", kosong)
	} else if bodyRequest.CardNumber == "4359650100001203" {
		httpCode, res = helpers.ResSuccess(false, "1005", "Card status expired", kosong)
	} else if bodyRequest.CardNumber == "4359650100001204" {
		httpCode, res = helpers.ResSuccess(false, "1006", "Card status blocked / inactive", kosong)
	} else if bodyRequest.PhoneNumber == "081208120812" {
		httpCode, res = helpers.ResSuccess(false, "1007", "Phone number does not match", kosong)
	} else if bodyRequest.CardNumber == "4359650100001205" {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if bodyRequest.CardNumber == "4359650100001206" {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		httpCode, res = usecase.OtpUsecase.GenerateOtpUsecase(kosong, idRequest, bodyRequest.JenisTransaksi, bodyRequest.PhoneNumber)
		if res.ResponseCode == "0000" {
			dataOutput := dto.ResDownstreamCheckCreditCard{
				Name:        "SATRIA PINAYUNGAN",
				PhoneNumber: bodyRequest.PhoneNumber,
				CardNumber:  helpers.Masking(bodyRequest.CardNumber),
			}
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		}
	}

	return
}
