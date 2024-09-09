package delivery_auth

import (
	"encoding/json"
	"io"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Login godoc
// @Summary Login
// @Description Login to get bearer token
// @Tags Login
// @Accept json
// @Produce json
// @Param Request body dto.ReqLogin true "Body Request"
// @Success 200 {object} dto.SwaggerLoginSuccess
// @Failure 401 {object} dto.SwaggerInvalidCredential
// @Failure 403 {object} dto.SwaggerIpBlocked
// @Failure 500 {object} dto.SwaggerGeneralSystemError
// @Router /api/v1/login [post]
func (delivery *AuthDeliveryImpl) LoginDelivery(ginContext *gin.Context) {
	idRequest, _, ipClient, transaksi, _, kosong := helpers.ConfigInit(ginContext)

	tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
	defer tx.End()
	tx.Context.SetLabel("idRequest", idRequest)

	var response dto.Res
	var httpCode int
	idUser := "0"

	request, err := io.ReadAll(ginContext.Request.Body)
	requestString := string(request)
	reqStruct := dto.ReqLogin{}
	json.Unmarshal(request, &reqStruct)
	if err != nil {
		httpCode, response = helpers.ResInvalidCredential(kosong)
	} else {
		err := delivery.Validate.Struct(reqStruct)
		if err != nil {
			httpCode, response = helpers.ResInvalidCredential(kosong)
		} else {
			httpCode, response, idUser = delivery.AuthUsecase.LoginUsecase(kosong, reqStruct)
		}
	}

	httpCode, response = delivery.AuthUsecase.ValidateBlockUsecase(ipClient, httpCode, response, kosong)

	ginContext.Set("user_id", idUser)
	activityLogParam := helpers.BuildActivityLogParam(idRequest, requestString, httpCode, kosong, ginContext, response)
	httpCode, response = delivery.LogUsecase.InsertLogActivityUsecase(activityLogParam, "")
	tx.Result = response.ResponseCode
	ginContext.JSON(httpCode, response)
}
