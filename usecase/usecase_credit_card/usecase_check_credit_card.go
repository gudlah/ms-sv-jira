package usecase_credit_card

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *CreditCardUsecaseImpl) CheckCreditCardUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckCreditCard) (httpCode int, res dto.Res) {
	reqUpstream := dto.ReqUpstreamInquiryCustomerCreditCard{
		Request: dto.RequestUpstreamInquiryCustomerCreditCard{
			CardNumber: bodyRequest.CardNumber,
			ChannelId:  usecase.BrigateConfig.ChannelId,
			ExternalId: helpers.GenerateRandomNumeric(12),
		},
	}

	reqUpstreamString, _ := json.Marshal(reqUpstream)
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   string(reqUpstreamString),
		RequestTimestamp: helpers.Now(),
		Service:          "ESB",
		Action:           "F68 - Inquiry Customer Credit Card",
	}

	resUpstream, err := usecase.ExternalRepository.InquiryCustomerCreditCardRepository(reqUpstream)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		resStruct := dto.ResUpstreamInquiryCustomerCreditCard{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		httpCode, res, logUpstream.IsSuccess = usecase.MappingOutputCheckCreditCard(kosong, bodyRequest, resStruct)

		if res.ResponseCode == "0000" {
			httpCode, res = usecase.OtpUsecase.GenerateOtpUsecase(kosong, idRequest, bodyRequest.JenisTransaksi, bodyRequest.PhoneNumber)
		}

		if res.ResponseCode == "0000" {
			dataOutput := dto.ResDownstreamCheckCreditCard{
				Name:        resStruct.Response.EmployerName,
				PhoneNumber: bodyRequest.PhoneNumber,
				CardNumber:  helpers.Masking(bodyRequest.CardNumber),
			}
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
		}

	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
