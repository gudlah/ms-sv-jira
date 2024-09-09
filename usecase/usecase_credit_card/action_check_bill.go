package usecase_credit_card

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *CreditCardUsecaseImpl) CheckBillAction(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckBill) (httpCode int, res dto.Res) {
	reqUpstream := dto.ReqUpstreamGetCreditCardMsr{
		Request: dto.RequestUpstreamGetCreditCardMsr{
			CardNumber: bodyRequest.CardNumber,
			ChannelId:  usecase.BrigateConfig.ChannelId,
			ExternalId: helpers.GenerateRandomNumeric(12),
			ServiceId:  usecase.BrigateConfig.GetCreditCardServiceId,
		},
	}

	reqUpstreamString, _ := json.Marshal(reqUpstream)
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   string(reqUpstreamString),
		RequestTimestamp: helpers.Now(),
		Service:          "ESB",
		Action:           "F28 - Get Credit Card MSR",
	}

	resUpstream, err := usecase.ExternalRepository.GetCreditCardMsrRepository(reqUpstream)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()

		resStruct := dto.ResUpstreamGetCreditCardMsr{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		httpCode, res, logUpstream.IsSuccess = usecase.MappingOutputCheckBill(kosong, bodyRequest, resStruct)
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
