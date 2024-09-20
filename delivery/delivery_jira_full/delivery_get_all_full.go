package delivery_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Get All Full godoc
// @Summary Get All Full
// @Description Get All Full
// @Tags Jira
// @Accept json
// @Produce json
// @param Authorization header string true "Authorization" default(Basic <Add access token here>)
// @Success 200 {object} dto.SwaggerGetAllFullSuccess
// @Failure 400 {object} dto.SwaggerInvalidValue
// @Failure 401 {object} dto.SwaggerInvalidCredential
// @Failure 403 {object} dto.SwaggerIpBlocked
// @Failure 500 {object} dto.SwaggerGeneralSystemError
// @Failure 502 {object} dto.SwaggerBackendError
// @Router /api/v1/all [get]
func (delivery *JiraFullDeliveryImpl) GetAllFullDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, _ := helpers.ConfigInit(ginContext)
	kosong := make([]dto.ResDownstreamGetAllProject, 0)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int

	httpCode, response = delivery.JiraFullUsecase.GetAllFullUsecase(kosong, idRequest)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam)
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
