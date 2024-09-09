package middleware

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/usecase/usecase_jwt"
	"ms-sv-jira/usecase/usecase_log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func JWTauth(logUsecase usecase_log.LogUsecase, jwtUsecase usecase_jwt.JWTUsecase, authDelivery delivery_auth.AuthDelivery) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		idRequest, _, ipClient, transaksi, _, kosong := helpers.ConfigInit(ginContext)

		tx := apm.DefaultTracer.StartTransaction(transaksi, "response_code")
		defer tx.End()
		tx.Context.SetLabel("idRequest", idRequest)

		var response dto.Res
		var httpCode int
		var err error
		idUser := "0"

		authHeader := ginContext.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			httpCode, response = helpers.ResInvalidCredential(kosong)
		} else {
			token := strings.Replace(authHeader, "Bearer ", "", -1)
			idUser, err = jwtUsecase.ValidateTokenUsecase(token)
			if err != nil {
				httpCode, response = helpers.ResInvalidCredential(kosong)
			} else {
				httpCode, response = helpers.ResSuccess(true, "0000", "Successfully", nil)
			}
		}

		httpCode, response = authDelivery.ValidateBlockDelivery(ipClient, httpCode, response, kosong)

		if httpCode == 200 {
			ginContext.Set("user_id", idUser)
		} else {
			ginContext.Set("user_id", idUser)
			activityLogParam := helpers.BuildActivityLogParam(idRequest, "", httpCode, kosong, ginContext, response)
			httpCode, response = logUsecase.InsertLogActivityUsecase(activityLogParam, "")
			tx.Result = response.ResponseCode
			ginContext.AbortWithStatusJSON(httpCode, response)
		}
	}
}
