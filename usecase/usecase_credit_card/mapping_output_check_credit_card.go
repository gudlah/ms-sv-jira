package usecase_credit_card

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *CreditCardUsecaseImpl) MappingOutputCheckCreditCard(kosong interface{}, bodyRequest dto.ReqDownstreamCheckCreditCard, resStruct dto.ResUpstreamInquiryCustomerCreditCard) (httpCode int, res dto.Res, isSuccess int) {
	resMapping, err := usecase.DatabaseRepository.GetMappingUpstreamResponse("/verif-credit-phone", resStruct.Response.ResponseCode)

	resSukses := false
	if resMapping.IsSuccess == 1 {
		resSukses = true
	}

	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if resMapping.ResponseCodeBribrain == "2002" {
		httpCode, res = helpers.ResBackendError(kosong)
	} else if resMapping.ResponseCodeBribrain == "0000" {
		phoneNumberResponse := helpers.TrimAll(resStruct.Response.HomePhone)
		if phoneNumberResponse != bodyRequest.PhoneNumber {
			httpCode, res = helpers.ResSuccess(resSukses, "1007", "Phone number does not match", kosong)
		} else {
			isSuccess = resMapping.IsSuccess
			dataOutput := dto.ResDownstreamCheckCreditCard{
				Name:        helpers.TrimAll(resStruct.Response.EmployerName),
				PhoneNumber: phoneNumberResponse,
				CardNumber:  helpers.Masking(bodyRequest.CardNumber),
			}
			httpCode, res = helpers.ResSuccess(resSukses, "0000", "Successfully", dataOutput)
		}
	} else if resMapping.Id == 0 {
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		httpCode, res = helpers.ResSuccess(resSukses, resMapping.ResponseCodeBribrain, resMapping.MessageBribrain, kosong)
	}

	return
}
