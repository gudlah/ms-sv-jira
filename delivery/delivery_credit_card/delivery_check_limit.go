package delivery_credit_card

import (
	"encoding/json"
	"io"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Check Limit godoc
// @Summary Check Limit
// @Description Check limit of credit card
// @Tags Credit Card
// @Accept json
// @Produce json
// @param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param Request body dto.ReqDownstreamCheckLimit true "Body Request"
// @Success 200 {object} dto.SwaggerCheckLimitSuccess
// @Failure 400 {object} dto.SwaggerInvalidValue
// @Failure 401 {object} dto.SwaggerInvalidCredential
// @Failure 403 {object} dto.SwaggerIpBlocked
// @Failure 429 {object} dto.SwaggerCheckLimitBillReached
// @Failure 500 {object} dto.SwaggerGeneralSystemError
// @Failure 502 {object} dto.SwaggerBackendError
// @Router /api/v1/limit [post]
func (delivery *CreditCardDeliveryImpl) CheckLimitDelivery(ginContext *gin.Context) {
	idRequest, _, _, transaksi, _, kosong := helpers.ConfigInit(ginContext)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	request, err := io.ReadAll(ginContext.Request.Body)
	requestString := string(request)
	reqStruct := dto.ReqDownstreamCheckLimit{}
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
			httpCode, res = delivery.CreditCardUsecase.CheckLimitUsecase(idRequest, kosong, reqStruct)
		}
	}

	httpCode, res = delivery.AuthDelivery.ValidateSpecialCharDelivery(kosong, reqStruct.SessionId, httpCode, res)

	activityLogParam := helpers.BuildActivityLogParam(idRequest, requestString, httpCode, kosong, ginContext, res)
	httpCode, res = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam, reqStruct.SessionId)
	tx.Result = res.ResponseCode
	ginContext.JSON(httpCode, res)
}
