package usecase_mock

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *MockUsecaseImpl) CheckLimitUsecase(kosong interface{}, bodyRequest dto.ReqDownstreamCheckLimit) (httpCode int, res dto.Res) {
	if bodyRequest.CardNumber == "4359650100001201" {
		httpCode, res = helpers.ResSuccess(false, "1003", "Data not found", kosong)
	} else if bodyRequest.CardNumber == "4359650100001202" {
		httpCode, res = helpers.ResSuccess(false, "1004", "Card status closed", kosong)
	} else if bodyRequest.CardNumber == "4359650100001203" {
		httpCode, res = helpers.ResSuccess(false, "1005", "Card status expired", kosong)
	} else if bodyRequest.CardNumber == "4359650100001204" {
		httpCode, res = helpers.ResSuccess(false, "1006", "Card status blocked / inactive", kosong)
	} else if bodyRequest.CardNumber == "4359650100001205" {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if bodyRequest.CardNumber == "4359650100001206" {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		paramVerifOtp := dto.ParamVerifOtp{
			PhoneNumber:    bodyRequest.PhoneNumber,
			OtpCode:        bodyRequest.OtpCode,
			JenisTransaksi: "limit",
		}
		httpCode, res = usecase.OtpUsecase.VerifOtpUsecase(kosong, paramVerifOtp)
		if res.ResponseCode == "0000" {
			dataOutput := dto.ResDownstreamCheckLimit{
				CreditLimit:              "9000000",
				CashLimit:                "0",
				CardNumber:               helpers.Masking(bodyRequest.CardNumber),
				AvailCredit:              "8619230",
				AvailCreditIndicator:     "+",
				AvailableCredit:          "AvailableCredit",
				AvailableCreditIndicator: "+",
				DailyCashLimit:           "0000000000",
				AvailCashRemain:          "0000014398",
				AvailCashRemainIndicator: "+",
			}
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		}
	}

	return
}
