package delivery_auth

import (
	"encoding/json"
	"io"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func (delivery *AuthDeliveryImpl) UnblockIpDelivery(ginContext *gin.Context) {
	idRequest, _, ipClient, transaksi, _, kosong := helpers.ConfigInit(ginContext)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int

	request, err := io.ReadAll(ginContext.Request.Body)
	requestString := string(request)
	reqStruct := dto.ReqUnblockIp{}
	json.Unmarshal(request, &reqStruct)

	if err != nil {
		httpCode, response = helpers.ResInvalidValue(kosong)
	} else {
		err := delivery.Validate.Struct(reqStruct)
		if err != nil {
			httpCode, response = helpers.ResInvalidValue(kosong)
		} else {
			httpCode, response = delivery.AuthUsecase.UnblockIpUsecase(reqStruct, kosong)
		}
	}

	httpCode, response = delivery.AuthUsecase.ValidateBlockUsecase(ipClient, httpCode, response, kosong)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, requestString, httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam, "")
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
