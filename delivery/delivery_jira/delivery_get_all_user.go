package delivery_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Get All User godoc
// @Summary Get All User
// @Description Get All User
// @Tags Jira
// @Accept json
// @Produce json
// @param Authorization header string true "Authorization" default(Basic <Add access token here>)
// @Success 200 {object} dto.SwaggerGetAllUserSuccess
// @Failure 400 {object} dto.SwaggerInvalidValue
// @Failure 401 {object} dto.SwaggerInvalidCredential
// @Failure 403 {object} dto.SwaggerIpBlocked
// @Failure 500 {object} dto.SwaggerGeneralSystemError
// @Failure 502 {object} dto.SwaggerBackendError
// @Router /api/v1/users [get]
func (delivery *JiraDeliveryImpl) GetAllUserDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, _ := helpers.ConfigInit(ginContext)
	kosong := make([]dto.ResDownstreamGetAllProject, 0)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int

	httpCode, response = delivery.JiraUsecase.GetAllUserUsecase(kosong, idRequest)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam)
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
