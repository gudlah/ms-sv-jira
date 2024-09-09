package delivery_auth

import (
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
)

type AuthDelivery interface {
	LoginDelivery(kosong interface{}, request dto.ReqLogin) (httpCode int, res dto.Res, idUser string)
	ValidateBlockDelivery(ipClient string, httpCode int, res dto.Res, kosong interface{}) (int, dto.Res)
	ValidateSpecialCharDelivery(kosong interface{}, input string, statusCode int, response dto.Res) (int, dto.Res)
	PageNotFoundDelivery(ginContext *gin.Context)
	UnblockIpDelivery(ginContext *gin.Context)
}
