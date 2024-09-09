package usecase_credit_card

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *CreditCardUsecaseImpl) MappingOutputCheckBill(kosong interface{}, bodyRequest dto.ReqDownstreamCheckBill, resStruct dto.ResUpstreamGetCreditCardMsr) (httpCode int, res dto.Res, isSuccess int) {
	resMapping, err := usecase.DatabaseRepository.GetMappingUpstreamResponse("/bill", resStruct.Response.ResponseCode)

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
		dataOutput := dto.ResDownstreamCheckBill{
			BillName:                  resStruct.Response.BillName,
			BillAmount:                resStruct.Response.BillAmount,
			BillMinAmount:             resStruct.Response.BillMinAmount,
			BillDueDate:               resStruct.Response.BillDueDate,
			CardNumber:                helpers.Masking(resStruct.Response.CardNo),
			StatementBalance:          resStruct.Response.StatementBalance,
			StatementBalanceIndicator: resStruct.Response.StatementBalanceIndicator,
			TotalDueAmount:            resStruct.Response.TotalDueAmount,
			TotalDueAmountIndicator:   resStruct.Response.TotalDueAmountIndicator,
		}
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	} else if resMapping.Id == 0 {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(resSukses, resMapping.ResponseCodeBribrain, resMapping.MessageBribrain, kosong)
	}

	return
}
