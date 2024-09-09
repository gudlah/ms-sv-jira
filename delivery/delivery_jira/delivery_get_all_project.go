package delivery_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func (delivery *JiraDeliveryImpl) GetAllProjectDelivery(ginContext *gin.Context) {
	idRequest, _, ipClient, transaksi, _, _ := helpers.ConfigInit(ginContext)
	kosong := make([]dto.ResUpstreamGetAllProject, 0)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int

	httpCode, response = delivery.JiraUsecase.GetAllProjectUsecase(kosong, idRequest)

	httpCode, response = delivery.AuthUsecase.ValidateBlockUsecase(ipClient, httpCode, response, kosong)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam)
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
