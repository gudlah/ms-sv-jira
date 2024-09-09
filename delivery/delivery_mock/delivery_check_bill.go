package delivery_mock

import (
	"encoding/json"
	"io"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func (delivery *MockDeliveryImpl) CheckBillDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, kosong := helpers.ConfigInit(ginContext)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	request, err := io.ReadAll(ginContext.Request.Body)
	requestString := string(request)
	reqStruct := dto.ReqDownstreamCheckBill{}
	json.Unmarshal(request, &reqStruct)
	var res dto.Res
	var httpCode int
	if err != nil {
		httpCode, res = helpers.ResInvalidValue(kosong)
	} else {
		err := delivery.Validate.Struct(reqStruct)
		if err != nil {
			httpCode, res = helpers.ResInvalidValue(kosong)
		} else {
			httpCode, res = delivery.MockUsecase.CheckBillUsecase(kosong, reqStruct)
		}
	}

	httpCode, res = delivery.AuthDelivery.ValidateSpecialCharDelivery(kosong, reqStruct.SessionId, httpCode, res)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, requestString, httpCode, kosong, ginContext, res)
	httpCode, res = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam, reqStruct.SessionId)
	tx.Result = res.ResponseCode
	ginContext.JSON(httpCode, res)
}
