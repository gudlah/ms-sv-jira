package delivery_jira

import (
	"encoding/json"
	"io"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Get Cards godoc
// @Summary Get Cards
// @Description Get Cards
// @Tags Jira
// @Accept json
// @Produce json
// @param Authorization header string true "Authorization" default(Basic <Add access token here>)
// @Param Request body dto.ReqDownstreamGetAllCard true "Body Request"
// @Success 200 {object} dto.SwaggerGetAlCardSuccess
// @Failure 400 {object} dto.SwaggerInvalidValue
// @Failure 401 {object} dto.SwaggerInvalidCredential
// @Failure 403 {object} dto.SwaggerIpBlocked
// @Failure 500 {object} dto.SwaggerGeneralSystemError
// @Failure 502 {object} dto.SwaggerBackendError
// @Router /api/v1/cards [post]
func (delivery *JiraDeliveryImpl) GetAllCardDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, _ := helpers.ConfigInit(ginContext)
	kosong := make([]dto.ReqDownstreamGetAllCard, 0)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	request, err := io.ReadAll(ginContext.Request.Body)
	requestString := string(request)
	reqStruct := dto.ReqDownstreamGetAllCard{}
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
			httpCode, res = delivery.JiraUsecase.GetAllCardUsecase(kosong, idRequest, reqStruct)
		}
	}

	activityLogParam := helpers.BuildActivityLogParam(idRequest, requestString, httpCode, kosong, ginContext, res)
	httpCode, res = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam)
	tx.Result = res.ResponseCode
	ginContext.JSON(httpCode, res)
}
