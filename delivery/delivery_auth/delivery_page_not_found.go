package delivery_auth

import (
	"ms-sv-jira/helpers"

	"github.com/gin-gonic/gin"
)

func (delivery *AuthDeliveryImpl) PageNotFoundDelivery(ginContext *gin.Context) {
	httpCode, res := helpers.ResPageNotFound()
	ginContext.JSON(httpCode, res)
}
