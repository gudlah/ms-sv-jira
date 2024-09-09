package usecase_credit_card

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *CreditCardUsecaseImpl) MappingOutputCheckLimit(kosong interface{}, bodyRequest dto.ReqDownstreamCheckLimit, resStruct dto.ResUpstreamGetCreditCardMsr) (httpCode int, res dto.Res, isSuccess int) {
	resMapping, err := usecase.DatabaseRepository.GetMappingUpstreamResponse("/limit", resStruct.Response.ResponseCode)

	resSukses := false
	if resMapping.IsSuccess == 1 {
		resSukses = true
	}

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if resMapping.ResponseCodeBribrain == "2002" {
		httpCode, res = helpers.ResBackendError(kosong)
	} else if resMapping.ResponseCodeBribrain == "0000" {
		isSuccess = resMapping.IsSuccess
		dataOutput := dto.ResDownstreamCheckLimit{
			CreditLimit:              resStruct.Response.CreditLimit,
			CashLimit:                resStruct.Response.CashLimit,
			CardNumber:               helpers.Masking(resStruct.Response.CardNo),
			AvailCredit:              resStruct.Response.AvailCredit,
			AvailCreditIndicator:     resStruct.Response.AvailCreditIndicator,
			AvailableCredit:          resStruct.Response.AvailableCredit,
			AvailableCreditIndicator: resStruct.Response.AvailableCreditIndicator,
			DailyCashLimit:           resStruct.Response.DailyCashLimit,
			AvailCashRemain:          resStruct.Response.AvailCashRemain,
			AvailCashRemainIndicator: resStruct.Response.AvailCashRemainIndicator,
		}
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	} else if resMapping.Id == 0 {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(resSukses, resMapping.ResponseCodeBribrain, resMapping.MessageBribrain, kosong)
	}

	return
}
