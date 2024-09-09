package middleware

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/usecase/usecase_log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func BasicAuth(authDelivery delivery_auth.AuthDelivery, logUsecase usecase_log.LogUsecase) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		idRequest, _, _, transaksi, _, kosong := helpers.ConfigInit(ginContext)

		tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
		defer tx.End()
		tx.Context.SetLabel("idRequest", idRequest)

		var idUser interface{}
		var response dto.Res
		var httpCode int

		username, password, ok := ginContext.Request.BasicAuth()

		if !ok {
			httpCode, response = helpers.ResInvalidCredential(kosong)
		} else {
			reqLogin := dto.ReqLogin{
				Username: username,
				Password: password,
			}
			httpCode, response, idUser = authDelivery.LoginDelivery(kosong, reqLogin)
			if response.ResponseCode != "0000" {
				httpCode, response = helpers.ResInvalidCredential(kosong)
			} else {
				httpCode, response = helpers.ResSuccess(true, "0000", "Successfully", kosong)
			}
		}

		ginContext.Set("idUser", idUser)

		if httpCode != 200 {
			activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
			httpCode, response = logUsecase.InsertLogActivityUsecase(activityLogParam)
			tx.Result = strconv.Itoa(httpCode)
			ginContext.AbortWithStatusJSON(httpCode, response)
			return
		}
	}
}
