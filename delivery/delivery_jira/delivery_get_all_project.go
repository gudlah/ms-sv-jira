package delivery_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func (delivery *JiraDeliveryImpl) GetAllProjectDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, _ := helpers.ConfigInit(ginContext)
	kosong := make([]dto.ResDownstreamGetAllProject, 0)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int

	httpCode, response = delivery.JiraUsecase.GetAllProjectUsecase(kosong, idRequest)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam)
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
